package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/halfrost/LeetCode-Go/ctl/util"
	"github.com/spf13/cobra"
)

var (
	chapterOneFileOrder = []string{"_index", "Data_Structure", "Algorithm", "Time_Complexity"}
	chapterOneMenuOrder = []string{"_index", "#关于作者", "Data_Structure", "Algorithm", "Time_Complexity"}
	chapterTwoFileOrder = []string{"_index", "Array", "String", "Two_Pointers", "Linked_List", "Stack", "Tree", "Dynamic_Programming", "Backtracking", "Depth_First_Search", "Breadth_First_Search",
		"Binary_Search", "Math", "Hash_Table", "Sorting", "Bit_Manipulation", "Union_Find", "Sliding_Window", "Segment_Tree", "Binary_Indexed_Tree"}
	chapterThreeFileOrder = []string{"_index", "Segment_Tree", "UnionFind", "LRUCache", "LFUCache"}
	preNextHeader         = "----------------------------------------------\n<div style=\"display: flex;justify-content: space-between;align-items: center;\">\n"
	preNextFotter         = "</div>"
	delLine               = "----------------------------------------------\n"
	delHeader             = "<div style=\"display: flex;justify-content: space-between;align-items: center;\">"
	delLabel              = "<[a-zA-Z]+.*?>([\\s\\S]*?)</[a-zA-Z]*?>"
	delFooter             = "</div>"

	//ErrNoFilename is thrown when the path the the file to tail was not given
	ErrNoFilename = errors.New("You must provide the path to a file in the \"-file\" flag.")

	//ErrInvalidLineCount is thrown when the user provided 0 (zero) as the value for number of lines to tail
	ErrInvalidLineCount = errors.New("You cannot tail zero lines.")

	chapterMap = map[string]map[string]string{
		"ChapterOne": {
			"_index":          "第一章 序章",
			"Data_Structure":  "1.1 数据结构知识",
			"Algorithm":       "1.2 算法知识",
			"Time_Complexity": "1.3 时间复杂度",
		},
		"ChapterTwo": {
			"_index":               "第二章 算法专题",
			"Array":                "2.01 Array",
			"String":               "2.02 String",
			"Two_Pointers":         "2.03 ✅ Two Pointers",
			"Linked_List":          "2.04 ✅ Linked List",
			"Stack":                "2.05 ✅ Stack",
			"Tree":                 "2.06 Tree",
			"Dynamic_Programming":  "2.07 Dynamic Programming",
			"Backtracking":         "2.08 ✅ Backtracking",
			"Depth_First_Search":   "2.09 Depth First Search",
			"Breadth_First_Search": "2.10 Breadth First Search",
			"Binary_Search":        "2.11 Binary Search",
			"Math":                 "2.12 Math",
			"Hash_Table":           "2.13 Hash Table",
			"Sorting":              "2.14 ✅ Sorting",
			"Bit_Manipulation":     "2.15 ✅ Bit Manipulation",
			"Union_Find":           "2.16 ✅ Union Find",
			"Sliding_Window":       "2.17 ✅ Sliding Window",
			"Segment_Tree":         "2.18 ✅ Segment Tree",
			"Binary_Indexed_Tree":  "2.19 ✅ Binary Indexed Tree",
		},
		"ChapterThree": {
			"_index":       "第三章 一些模板",
			"Segment_Tree": "3.1 Segment Tree",
			"UnionFind":    "3.2 UnionFind",
			"LRUCache":     "3.3 LRUCache",
			"LFUCache":     "3.4 LFUCache",
		},
		"ChapterFour": {
			"_index": "第四章 Leetcode 题解",
		},
	}
)

func getChapterFourFileOrder() ([]string, []int) {
	solutions, solutionIds := util.LoadChapterFourDir()
	chapterFourFileOrder := []string{"_index"}
	chapterFourFileOrder = append(chapterFourFileOrder, solutions...)
	fmt.Printf("ChapterFour 中包括 _index 有 %v 个文件, len(id) = %v\n", len(chapterFourFileOrder), len(solutionIds))
	return chapterFourFileOrder, solutionIds
}

func newLabelCommand() *cobra.Command {
	mc := &cobra.Command{
		Use:   "label <subcommand>",
		Short: "Label related commands",
	}
	//mc.PersistentFlags().StringVar(&logicEndpoint, "endpoint", "localhost:5880", "endpoint of logic service")
	mc.AddCommand(
		newAddPreNext(),
		newDeletePreNext(),
	)
	return mc
}

func newAddPreNext() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-pre-next",
		Short: "Add pre-next label",
		Run: func(cmd *cobra.Command, args []string) {
			addPreNext()
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func newDeletePreNext() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del-pre-next",
		Short: "Delete pre-next label",
		Run: func(cmd *cobra.Command, args []string) {
			delPreNext()
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func addPreNext() {
	// Chpater one add pre-next
	addPreNextLabel(chapterOneFileOrder, []string{}, []int{}, "", "ChapterOne", "ChapterTwo")
	// Chpater two add pre-next
	addPreNextLabel(chapterTwoFileOrder, chapterOneFileOrder, []int{}, "ChapterOne", "ChapterTwo", "ChapterThree")
	// Chpater three add pre-next
	addPreNextLabel(chapterThreeFileOrder, chapterTwoFileOrder, []int{}, "ChapterTwo", "ChapterThree", "ChapterFour")
	// Chpater four add pre-next
	//fmt.Printf("%v\n", getChapterFourFileOrder())
	chapterFourFileOrder, solutionIds := getChapterFourFileOrder()
	addPreNextLabel(chapterFourFileOrder, chapterThreeFileOrder, solutionIds, "ChapterThree", "ChapterFour", "")
}

func addPreNextLabel(order, preOrder []string, chapterFourIds []int, preChapter, chapter, nextChapter string) {
	var (
		exist bool
		err   error
		res   []byte
		count int
	)
	for index, path := range order {
		tmp := ""
		if index == 0 {
			if chapter == "ChapterOne" {
				// 第一页不需要“上一章”
				tmp = "\n\n" + delLine + fmt.Sprintf("<p align = \"right\"><a href=\"https://books.halfrost.com/leetcode/%v/%v/\">下一页➡️</a></p>\n", chapter, order[index+1])
			} else {
				if chapter == "ChapterFour" {
					tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/\">⬅️上一章</a></p>\n", preChapter, preOrder[len(preOrder)-1]) + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/%v/\">下一页➡️</a></p>\n", chapter, util.GetChpaterFourFileNum(chapterFourIds[(index-1)+1]), order[index+1]) + preNextFotter
				} else {
					tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/\">⬅️上一章</a></p>\n", preChapter, preOrder[len(preOrder)-1]) + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/\">下一页➡️</a></p>\n", chapter, order[index+1]) + preNextFotter
				}
			}
		} else if index == len(order)-1 {
			if chapter == "ChapterFour" {
				// 最后一页不需要“下一页”
				tmp = "\n\n" + delLine + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/%v/\">⬅️上一页</a></p>\n", chapter, util.GetChpaterFourFileNum(chapterFourIds[(index-1)-1]), order[index-1])
			} else {
				tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/\">⬅️上一页</a></p>\n", chapter, order[index-1]) + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/\">下一章➡️</a></p>\n", nextChapter) + preNextFotter
			}
		} else if index == 1 {
			if chapter == "ChapterFour" {
				tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/\">⬅️上一页</a></p>\n", chapter) + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/%v/\">下一页➡️</a></p>\n", chapter, util.GetChpaterFourFileNum(chapterFourIds[(index-1)+1]), order[index+1]) + preNextFotter
			} else {
				tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/\">⬅️上一页</a></p>\n", chapter) + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/\">下一页➡️</a></p>\n", chapter, order[index+1]) + preNextFotter
			}
		} else {
			if chapter == "ChapterFour" {
				tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/%v/\">⬅️上一页</a></p>\n", chapter, util.GetChpaterFourFileNum(chapterFourIds[(index-1)-1]), order[index-1]) + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/%v/\">下一页➡️</a></p>\n", chapter, util.GetChpaterFourFileNum(chapterFourIds[(index-1)+1]), order[index+1]) + preNextFotter
			} else {
				tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/\">⬅️上一页</a></p>\n", chapter, order[index-1]) + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/%v/%v/\">下一页➡️</a></p>\n", chapter, order[index+1]) + preNextFotter
			}
		}

		if chapter == "ChapterFour" && index > 0 {
			path = fmt.Sprintf("%v/%v", util.GetChpaterFourFileNum(chapterFourIds[(index-1)]), path)
		}

		exist, err = needAdd(fmt.Sprintf("../website/content/%v/%v.md", chapter, path))
		if err != nil {
			fmt.Println(err)
			return
		}
		// 当前没有上一页和下一页，才添加
		if !exist && err == nil {
			res, err = eofAdd(fmt.Sprintf("../website/content/%v/%v.md", chapter, path), tmp)
			if err != nil {
				fmt.Println(err)
				return
			}
			util.WriteFile(fmt.Sprintf("../website/content/%v/%v.md", chapter, path), res)
			count++
		}
	}
	fmt.Printf("添加了 %v 个文件的 pre-next\n", count)
}

func eofAdd(filePath string, labelString string) ([]byte, error) {
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
				output = append(output, []byte(labelString)...)
				output = append(output, []byte("\n")...)
				return output, nil
			}
			return nil, err
		}
		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}
}

func delPreNext() {
	// Chpater one del pre-next
	delPreNextLabel(chapterOneFileOrder, []int{}, "ChapterOne")
	// Chpater two del pre-next
	delPreNextLabel(chapterTwoFileOrder, []int{}, "ChapterTwo")
	// Chpater three del pre-next
	delPreNextLabel(chapterThreeFileOrder, []int{}, "ChapterThree")
	// Chpater four del pre-next
	chapterFourFileOrder, solutionIds := getChapterFourFileOrder()
	delPreNextLabel(chapterFourFileOrder, solutionIds, "ChapterFour")
}

func delPreNextLabel(order []string, chapterFourIds []int, chapter string) {
	count := 0
	for index, path := range order {
		lineNum := 5
		if index == 0 && chapter == "ChapterOne" || index == len(order)-1 && chapter == "ChapterFour" {
			lineNum = 3
		}
		if chapter == "ChapterFour" && index > 0 {
			path = fmt.Sprintf("%v/%v", util.GetChpaterFourFileNum(chapterFourIds[(index-1)]), path)
		}

		exist, err := needAdd(fmt.Sprintf("../website/content/%v/%v.md", chapter, path))
		if err != nil {
			fmt.Println(err)
			return
		}
		// 存在才删除
		if exist && err == nil {
			removeLine(fmt.Sprintf("../website/content/%v/%v.md", chapter, path), lineNum+1)
			count++
		}
	}
	fmt.Printf("删除了 %v 个文件的 pre-next\n", count)
	// 另外一种删除方法
	// res, err := eofDel(fmt.Sprintf("../website/content/ChapterOne/%v.md", v))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// util.WriteFile(fmt.Sprintf("../website/content/ChapterOne/%v.md", v), res)
}

func needAdd(filePath string) (bool, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()
	reader, output := bufio.NewReader(f), []byte{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return false, nil
			}
			return false, err
		}
		if ok, _ := regexp.Match(delHeader, line); ok {
			return true, nil
		} else if ok, _ := regexp.Match(delLabel, line); ok {
			return true, nil
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
}

func eofDel(filePath string) ([]byte, error) {
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
		if ok, _ := regexp.Match(delLine, line); ok {
			reg := regexp.MustCompile(delLine)
			newByte := reg.ReplaceAll(line, []byte(""))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match(delHeader, line); ok {
			reg := regexp.MustCompile(delHeader)
			newByte := reg.ReplaceAll(line, []byte(""))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match(delLabel, line); ok {
			reg := regexp.MustCompile(delLabel)
			newByte := reg.ReplaceAll(line, []byte(""))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match(delFooter, line); ok {
			reg := regexp.MustCompile(delFooter)
			newByte := reg.ReplaceAll(line, []byte(""))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
}

func removeLine(path string, lineNumber int) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	info, _ := os.Stat(path)
	mode := info.Mode()
	array := strings.Split(string(file), "\n")
	array = array[:len(array)-lineNumber-1]
	ioutil.WriteFile(path, []byte(strings.Join(array, "\n")), mode)
	//fmt.Println("remove line successful")
}
