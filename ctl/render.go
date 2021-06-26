package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	m "github.com/halfrost/LeetCode-Go/ctl/models"
	"github.com/halfrost/LeetCode-Go/ctl/util"
	"github.com/spf13/cobra"
)

var (
	chapterTwoList = []string{"Array", "String", "Two Pointers", "Linked List", "Stack", "Tree", "Dynamic Programming", "Backtracking", "Depth First Search", "Breadth First Search",
		"Binary Search", "Math", "Hash Table", "Sorting", "Bit Manipulation", "Union Find", "Sliding Window", "Segment Tree", "Binary Indexed Tree"}
	chapterTwoFileName = []string{"Array", "String", "Two_Pointers", "Linked_List", "Stack", "Tree", "Dynamic_Programming", "Backtracking", "Depth_First_Search", "Breadth_First_Search",
		"Binary_Search", "Math", "Hash_Table", "Sorting", "Bit_Manipulation", "Union_Find", "Sliding_Window", "Segment_Tree", "Binary_Indexed_Tree"}
	chapterTwoSlug = []string{"array", "string", "two-pointers", "linked-list", "stack", "tree", "dynamic-programming", "backtracking", "depth-first-search", "breadth-first-search",
		"binary-search", "math", "hash-table", "sorting", "bit-manipulation", "union-find", "sliding-window", "segment-tree", "binary-indexed-tree"}
)

func newBuildCommand() *cobra.Command {
	mc := &cobra.Command{
		Use:   "build <subcommand>",
		Short: "Build doc related commands",
	}
	//mc.PersistentFlags().StringVar(&logicEndpoint, "endpoint", "localhost:5880", "endpoint of logic service")
	mc.AddCommand(
		newBuildREADME(),
		newBuildChapterTwo(),
		// newBuildMenu(),
	)
	return mc
}

func newBuildREADME() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "readme",
		Short: "Build readme.md commands",
		Run: func(cmd *cobra.Command, args []string) {
			buildREADME()
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func newBuildChapterTwo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chapter-two",
		Short: "Build Chapter Two commands",
		Run: func(cmd *cobra.Command, args []string) {
			buildChapterTwo(true)
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func newBuildMenu() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "menu",
		Short: "Build Menu commands",
		Run: func(cmd *cobra.Command, args []string) {
			buildBookMenu()
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func buildREADME() {
	var (
		problems []m.StatStatusPairs
		lpa      m.LeetCodeProblemAll
		info     m.UserInfo
	)
	// 请求所有题目信息
	body := getProblemAllList()
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
	mdrows := m.ConvertMdModelFromSsp(problems)
	sort.Sort(m.SortByQuestionID(mdrows))
	solutionIds, _, try := util.LoadSolutionsDir()
	m.GenerateMdRows(solutionIds, mdrows)
	info.EasyTotal, info.MediumTotal, info.HardTotal, info.OptimizingEasy, info.OptimizingMedium, info.OptimizingHard, optimizingIds = statisticalData(problemsMap, solutionIds)
	omdrows := m.ConvertMdModelFromIds(problemsMap, optimizingIds)
	sort.Sort(m.SortByQuestionID(omdrows))

	// 按照模板渲染 README
	res, err := renderReadme("./template/template.markdown", len(solutionIds), try, m.Mdrows{Mdrows: mdrows}, m.Mdrows{Mdrows: omdrows}, info)
	if err != nil {
		fmt.Println(err)
		return
	}
	util.WriteFile("../README.md", res)
	fmt.Println("write file successful")
	//makeReadmeFile(mds)
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

// internal: true  渲染的链接都是 hugo 内部链接，用户生成 hugo web
// 			 false 渲染的链接是外部 HTTPS 链接，用于生成 PDF
func buildChapterTwo(internal bool) {
	var (
		gr        m.GraphQLResp
		questions []m.Question
		count     int
	)
	for index, tag := range chapterTwoSlug {
		body := getTagProblemList(tag)
		// fmt.Printf("%v\n", string(body))
		err := json.Unmarshal(body, &gr)
		if err != nil {
			fmt.Println(err)
			return
		}
		questions = gr.Data.TopicTag.Questions
		mdrows := m.ConvertMdModelFromQuestions(questions)
		sort.Sort(m.SortByQuestionID(mdrows))
		solutionIds, _, _ := util.LoadSolutionsDir()
		tl, err := loadMetaData(fmt.Sprintf("./meta/%v", chapterTwoFileName[index]))
		if err != nil {
			fmt.Printf("err = %v\n", err)
		}
		tls := m.GenerateTagMdRows(solutionIds, tl, mdrows, internal)
		//fmt.Printf("tls = %v\n", tls)
		//  按照模板渲染 README
		res, err := renderChapterTwo(fmt.Sprintf("./template/%v.md", chapterTwoFileName[index]), m.TagLists{TagLists: tls})
		if err != nil {
			fmt.Println(err)
			return
		}
		if internal {
			util.WriteFile(fmt.Sprintf("../website/content/ChapterTwo/%v.md", chapterTwoFileName[index]), res)
		} else {
			util.WriteFile(fmt.Sprintf("./pdftemp/ChapterTwo/%v.md", chapterTwoFileName[index]), res)
		}

		count++
	}
	fmt.Printf("write %v files successful", count)
}

func loadMetaData(filePath string) (map[int]m.TagList, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader, metaMap := bufio.NewReader(f), map[int]m.TagList{}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return metaMap, nil
			}
			return nil, err
		}
		s := strings.Split(string(line), "|")
		v, _ := strconv.Atoi(strings.Split(s[1], ".")[0])
		// v[0] 是题号，s[4] time, s[5] space, s[6] favorite
		metaMap[v] = m.TagList{
			FrontendQuestionID: int32(v),
			Acceptance:         "",
			Difficulty:         "",
			TimeComplexity:     s[4],
			SpaceComplexity:    s[5],
			Favorite:           s[6],
		}
	}
}

func renderChapterTwo(filePath string, tls m.TagLists) ([]byte, error) {
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
		if ok, _ := regexp.Match("{{.AvailableTagTable}}", line); ok {
			reg := regexp.MustCompile("{{.AvailableTagTable}}")
			newByte := reg.ReplaceAll(line, []byte(tls.AvailableTagTable()))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
}

func buildBookMenu() {
	copyLackFile()
	// 按照模板重新渲染 Menu
	res, err := renderBookMenu("./template/menu.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	util.WriteFile("../website/content/menu/index.md", res)
	fmt.Println("generate Menu successful")
}

// 拷贝 leetcode 目录下的题解 README 文件至第四章对应文件夹中
func copyLackFile() {
	solutionIds, soName, _ := util.LoadSolutionsDir()
	_, ch4Ids := util.LoadChapterFourDir()

	needCopy := []string{}
	for i := 0; i < len(solutionIds); i++ {
		if util.BinarySearch(ch4Ids, solutionIds[i]) == -1 {
			needCopy = append(needCopy, soName[i])
		}
	}
	if len(needCopy) > 0 {
		fmt.Printf("有 %v 道题需要拷贝到第四章中\n", len(needCopy))
		for i := 0; i < len(needCopy); i++ {
			if needCopy[i][4] == '.' {
				tmp, err := strconv.Atoi(needCopy[i][:4])
				if err != nil {
					fmt.Println(err)
				}
				err = os.MkdirAll(fmt.Sprintf("../website/content/ChapterFour/%v", util.GetChpaterFourFileNum(tmp)), os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
				util.CopyFile(fmt.Sprintf("../website/content/ChapterFour/%v/%v.md", util.GetChpaterFourFileNum(tmp), needCopy[i]), fmt.Sprintf("../leetcode/%v/README.md", needCopy[i]))
				util.CopyFile(fmt.Sprintf("../website/content/ChapterFour/%v/_index.md", util.GetChpaterFourFileNum(tmp)), "./template/collapseSection.md")
			}
		}
	} else {
		fmt.Printf("【第四章没有需要添加的题解，已经完整了】\n")
	}
}

func generateMenu() string {
	res := ""
	res += menuLine(chapterOneMenuOrder, "ChapterOne")
	res += menuLine(chapterTwoFileOrder, "ChapterTwo")
	res += menuLine(chapterThreeFileOrder, "ChapterThree")
	chapterFourFileOrder, _ := getChapterFourFileOrder()
	res += menuLine(chapterFourFileOrder, "ChapterFour")
	return res
}

func menuLine(order []string, chapter string) string {
	res := ""
	for i := 0; i < len(order); i++ {
		if i == 1 && chapter == "ChapterOne" {
			res += fmt.Sprintf("  - [%v]({{< relref \"/%v/%v\" >}})\n", chapterMap[chapter][order[i]], chapter, order[i])
			continue
		}
		if i == 0 {
			res += fmt.Sprintf("- [%v]({{< relref \"/%v/%v.md\" >}})\n", chapterMap[chapter][order[i]], chapter, order[i])
		} else {
			if chapter == "ChapterFour" {
				res += fmt.Sprintf("    - [%v]({{< relref \"/%v/%v.md\" >}})\n", order[i], chapter, order[i])
			} else {
				res += fmt.Sprintf("  - [%v]({{< relref \"/%v/%v.md\" >}})\n", chapterMap[chapter][order[i]], chapter, order[i])
			}
		}
	}
	return res
}

func renderBookMenu(filePath string) ([]byte, error) {
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
		if ok, _ := regexp.Match("{{.BookMenu}}", line); ok {
			reg := regexp.MustCompile("{{.BookMenu}}")
			newByte := reg.ReplaceAll(line, []byte(generateMenu()))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
}
