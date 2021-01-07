package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	m "github.com/halfrost/LeetCode-Go/automation/models"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		result []m.StatStatusPairs
		lpa    m.LeetCodeProblemAll
	)

	body := getProblemAllList()
	err := json.Unmarshal(body, &lpa)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = lpa.StatStatusPairs
	mdrows := []m.Mdrow{}
	for i := 0; i < len(result); i++ {
		mdrows = append(mdrows, convertModel(result[i]))
	}
	sort.Sort(m.SortByQuestionId(mdrows))
	solutionIds := loadSolutionsDir()
	generateMdRows(solutionIds, mdrows)
	// res, _ := json.Marshal(mdrows)
	//writeFile("leetcode_problem", res)
	mds := m.Mdrows{Mdrows: mdrows}
	res, err := readFile("./template.markdown", "{{.AvailableTable}}", mds)
	if err != nil {
		fmt.Println(err)
		return
	}
	writeFile("../README.md", res)
	//makeReadmeFile(mds)
}

func getProblemAllList() []byte {
	resp, err := http.Get("https://leetcode.com/api/problems/all/")
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	return body
}

func writeFile(fileName string, content []byte) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write file successful")
}

func convertModel(ssp m.StatStatusPairs) m.Mdrow {
	res := m.Mdrow{}
	res.FrontendQuestionId = ssp.Stat.FrontendQuestionId
	res.QuestionTitle = ssp.Stat.QuestionTitle
	res.QuestionTitleSlug = ssp.Stat.QuestionTitleSlug
	res.Acceptance = fmt.Sprintf("%.1f%%", (ssp.Stat.TotalAcs/ssp.Stat.TotalSubmitted)*100)
	res.Difficulty = m.DifficultyMap[ssp.Difficulty.Level]
	res.Frequency = fmt.Sprintf("%f", ssp.Frequency)
	return res
}

func loadSolutionsDir() []int {
	files, err := ioutil.ReadDir("../leetcode/")
	if err != nil {
		fmt.Println(err)
	}
	solutionIds := []int{}
	for _, f := range files {
		if f.Name()[4] == '.' {
			tmp, err := strconv.Atoi(f.Name()[:4])
			if err != nil {
				fmt.Println(err)
			}
			solutionIds = append(solutionIds, tmp)
		}
	}
	sort.Ints(solutionIds)
	fmt.Printf("读取了 %v 道题的题解，当前目录下有 %v 个文件(可能包含 .DS_Store)，有 %v 道题在尝试中\n", len(solutionIds), len(files), len(files)-len(solutionIds))
	return solutionIds
}

func generateMdRows(solutionIds []int, mdrows []m.Mdrow) {
	for i := 0; i < len(solutionIds); i++ {
		id := mdrows[solutionIds[i]-1].FrontendQuestionId
		if solutionIds[i] == int(id) {
			//fmt.Printf("id = %v i = %v solutionIds = %v\n", id, i, solutionIds[i])
			mdrows[id-1].SolutionPath = fmt.Sprintf("[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/%v)", fmt.Sprintf("%04d.%v", id, strings.Replace(mdrows[id-1].QuestionTitle, " ", "-", -1)))
		} else {
			fmt.Printf("序号出错了 solutionIds = %v id = %v\n", solutionIds[i], id)
		}
	}
	fmt.Printf("")
}

func makeReadmeFile(mdrows m.Mdrows) {
	file := "./README.md"
	os.Remove(file)
	var b bytes.Buffer
	tmpl := template.Must(template.New("readme").Parse(readTMPL("template.markdown")))
	err := tmpl.Execute(&b, mdrows)
	if err != nil {
		fmt.Println(err)
	}
	// 保存 README.md 文件
	writeFile(file, b.Bytes())
}

func readTMPL(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func readFile(filePath, template string, mdrows m.Mdrows) ([]byte, error) {
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
		if ok, _ := regexp.Match(template, line); ok {
			reg := regexp.MustCompile(template)
			newByte := reg.ReplaceAll(line, []byte(mdrows.AvailableTable()))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
	return output, nil
}
