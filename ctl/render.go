package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	m "github.com/halfrost/LeetCode-Go/ctl/models"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var (
		problems []m.StatStatusPairs
		lpa      m.LeetCodeProblemAll
		info     m.UserInfo
	)
	// 请求所有题目信息
	body := getProblemAllList(AllProblemURL)
	problemsMap, optimizingIds := map[int]m.StatStatusPairs{}, []int{}
	err := json.Unmarshal(body, &lpa)
	if err != nil {
		fmt.Println(err)
		return
	}
	//writeFile("leetcode_problem.json", body)

	// 拼凑 README 需要渲染的数据
	problems = lpa.StatStatusPairs
	info = m.ConvertUserInfoModel(lpa)
	for _, v := range problems {
		problemsMap[int(v.Stat.FrontendQuestionID)] = v
	}
	mdrows := m.ConvertMdModel(problems)
	sort.Sort(m.SortByQuestionID(mdrows))
	solutionIds, try := loadSolutionsDir()
	generateMdRows(solutionIds, mdrows)
	info.EasyTotal, info.MediumTotal, info.HardTotal, info.OptimizingEasy, info.OptimizingMedium, info.OptimizingHard, optimizingIds = statisticalData(problemsMap, solutionIds)
	omdrows := m.ConvertMdModelFromIds(problemsMap, optimizingIds)
	sort.Sort(m.SortByQuestionID(omdrows))

	// 按照模板渲染 README
	res, err := renderReadme("./template.markdown", len(solutionIds), try, m.Mdrows{Mdrows: mdrows}, m.Mdrows{Mdrows: omdrows}, info)
	if err != nil {
		fmt.Println(err)
		return
	}
	writeFile("../README.md", res)
	//makeReadmeFile(mds)
}

func generateMdRows(solutionIds []int, mdrows []m.Mdrow) {
	for i := 0; i < len(solutionIds); i++ {
		id := mdrows[solutionIds[i]-1].FrontendQuestionID
		if solutionIds[i] == int(id) {
			//fmt.Printf("id = %v i = %v solutionIds = %v\n", id, i, solutionIds[i])
			mdrows[id-1].SolutionPath = fmt.Sprintf("[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/%v)", fmt.Sprintf("%04d.%v", id, strings.Replace(mdrows[id-1].QuestionTitle, " ", "-", -1)))
		} else {
			fmt.Printf("序号出错了 solutionIds = %v id = %v\n", solutionIds[i], id)
		}
	}
}

func renderReadme(filePath string, total, try int, mdrows, omdrows m.Mdrows, user m.UserInfo) ([]byte, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader, output := bufio.NewReader(f), []byte{}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, nil
			}
			return nil, err
		}
		if ok, _ := regexp.Match("{{.AvailableTable}}", line); ok {
			reg := regexp.MustCompile("{{.AvailableTable}}")
			newByte := reg.ReplaceAll(line, []byte(mdrows.AvailableTable()))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match("{{.TotalNum}}", line); ok {
			reg := regexp.MustCompile("{{.TotalNum}}")
			newByte := reg.ReplaceAll(line, []byte(fmt.Sprintf("以下已经收录了 %v 道题的题解，还有 %v 道题在尝试优化到 beats 100%%", total, try)))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match("{{.PersonalData}}", line); ok {
			reg := regexp.MustCompile("{{.PersonalData}}")
			newByte := reg.ReplaceAll(line, []byte(user.PersonalData()))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match("{{.OptimizingTable}}", line); ok {
			reg := regexp.MustCompile("{{.OptimizingTable}}")
			newByte := reg.ReplaceAll(line, []byte(fmt.Sprintf("以下 %v 道题还需要优化到 100%% 的题目列表\n\n%v", (user.OptimizingEasy+user.OptimizingMedium+user.OptimizingHard), omdrows.AvailableTable())))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
}
