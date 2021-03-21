package main

import (
	"bufio"
	"fmt"
	"github.com/halfrost/LeetCode-Go/ctl/util"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	cleanString1 = "{{< columns >}}"
	cleanString2 = "<--->"
	cleanString3 = "{{< /columns >}}"
	cleanString4 = "<img src=\"https://books.halfrost.com/leetcode/logo.png\" alt=\"logo\" height=\"600\" align=\"right\" style=\"padding-left: 30px;\"/>"
	pdfPreface   = `<img src="https://books.halfrost.com/leetcode/logo.png" alt="logo" heigth="1300px" align="center"/>


# 说明
	
此版本是 https://books.halfrost.com/leetcode 网页的离线版，由于网页版实时会更新，所以此 PDF 版难免会有一些排版或者错别字。如果读者遇到了，可以到网页版相应页面，点击页面 edit 按钮，提交 pr 进行更改。此 PDF 版本号是 V%v.%v.%v。PDF 永久更新地址是 https://github.com/halfrost/LeetCode-Go/releases/，以版本号区分不同版本。笔者还是强烈推荐看在线版，有任何错误都会立即更新。如果觉得此书对刷题有一点点帮助，可以给此书点一个 star，鼓励一下笔者早点更新更多题解。
	
> 版本号说明，V%v.%v.%v，%v 是大版本号，%v 代表当前题解中有几百题，目前是 %v 题，所以第二个版本号是 %v，%v 代表当前题解中有几十题，目前是 %v 题，所以第三个版本号是 %v 。
	
# 目录
	
[toc]

`

	majorVersion   = 1
	midVersion     = 0
	lastVersion    = 0
	totalSolutions = 0
)

func newPDFCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pdf",
		Short: "PDF related commands",
		Run: func(cmd *cobra.Command, args []string) {
			generatePDF()
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func generatePDF() {
	var (
		pdf, tmp []byte
		err      error
	)
	// 先删除 pre-next
	delPreNext()

	chapterFourFileOrder, _ := util.LoadChapterFourDir()
	totalSolutions = len(chapterFourFileOrder)
	midVersion = totalSolutions / 100
	lastVersion = totalSolutions % 100
	fmt.Printf("[当前的版本号是 V%v.%v.%v]\n", majorVersion, midVersion, lastVersion)
	// 删除原始文档中的头部，并创建临时文件夹
	prepare(fmt.Sprintf("../PDF v%v.%v.%v.md", majorVersion, midVersion, lastVersion))
	// PDF 前言
	pdf = append(pdf, []byte(fmt.Sprintf(pdfPreface, majorVersion, midVersion, lastVersion, majorVersion, midVersion, lastVersion, majorVersion, midVersion, totalSolutions, midVersion, lastVersion, totalSolutions, lastVersion))...)
	// PDF 第一章
	tmp, err = loadChapter(chapterOneFileOrder, "./pdftemp", "ChapterOne")
	pdf = append(pdf, tmp...)
	// PDF 第二章
	tmp, err = loadChapter(chapterTwoFileOrder, "./pdftemp", "ChapterTwo")
	pdf = append(pdf, tmp...)
	// PDF 第三章
	tmp, err = loadChapter(chapterThreeFileOrder, "./pdftemp", "ChapterThree")
	pdf = append(pdf, tmp...)
	// PDF 第四章
	tmp, err = util.LoadFile("./pdftemp/ChapterFour/_index.md")
	pdf = append(pdf, tmp...)
	tmp, err = loadChapter(chapterFourFileOrder, "../website/content", "ChapterFour")
	pdf = append(pdf, tmp...)
	if err != nil {
		fmt.Println(err)
	}
	// 生成 PDF
	util.WriteFile(fmt.Sprintf("../PDF v%v.%v.%v.md", majorVersion, midVersion, lastVersion), pdf)
	// 还原现场
	addPreNext()
	util.DestoryDir("./pdftemp")
}

func loadChapter(order []string, path, chapter string) ([]byte, error) {
	var (
		res, tmp []byte
		err      error
	)
	for index, v := range order {
		if chapter == "ChapterOne" && index == 0 {
			// 清理不支持的特殊 MarkDown 语法
			tmp, err = clean(fmt.Sprintf("%v/%v/%v.md", path, chapter, v))
		} else {
			if chapter == "ChapterFour" {
				if v[4] == '.' {
					num, err := strconv.Atoi(v[:4])
					if err != nil {
						fmt.Println(err)
					}
					tmp, err = util.LoadFile(fmt.Sprintf("%v/%v/%v/%v.md", path, chapter, util.GetChpaterFourFileNum(num), v))
				}
			} else {
				tmp, err = util.LoadFile(fmt.Sprintf("%v/%v/%v.md", path, chapter, v))
			}
		}
		if err != nil {
			fmt.Println(err)
			return []byte{}, err
		}
		res = append(res, tmp...)
	}
	return res, err
}

func prepare(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println("pdf 还没有创建")
		fmt.Println(err)
	}
	fmt.Println("pdf 删除成功，开始构建全新版本")

	err = os.MkdirAll("./pdftemp/ChapterOne", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range chapterOneFileOrder {
		removeHeader(fmt.Sprintf("../website/content/ChapterOne/%v.md", v), fmt.Sprintf("./pdftemp/ChapterOne/%v.md", v), 5)
	}

	err = os.MkdirAll("./pdftemp/ChapterTwo", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	// 生成外部链接的 ChapterTwo
	buildChapterTwo(false)
	util.CopyFile("./pdftemp/ChapterTwo/_index.md", "../website/content/ChapterTwo/_index.md")

	for _, v := range chapterTwoFileOrder {
		removeHeader(fmt.Sprintf("./pdftemp/ChapterTwo/%v.md", v), fmt.Sprintf("./pdftemp/ChapterTwo/%v.md", v), 5)
	}

	err = os.MkdirAll("./pdftemp/ChapterThree", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range chapterThreeFileOrder {
		removeHeader(fmt.Sprintf("../website/content/ChapterThree/%v.md", v), fmt.Sprintf("./pdftemp/ChapterThree/%v.md", v), 5)
	}

	err = os.MkdirAll("./pdftemp/ChapterFour", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	removeHeader(fmt.Sprintf("../website/content/ChapterFour/_index.md"), fmt.Sprintf("./pdftemp/ChapterFour/_index.md"), 5)
}

func clean(filePath string) ([]byte, error) {
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
		if ok, _ := regexp.Match(cleanString1, line); ok {
			reg := regexp.MustCompile(cleanString1)
			newByte := reg.ReplaceAll(line, []byte(""))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match(cleanString2, line); ok {
			reg := regexp.MustCompile(cleanString2)
			newByte := reg.ReplaceAll(line, []byte(""))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match(cleanString3, line); ok {
			reg := regexp.MustCompile(cleanString3)
			newByte := reg.ReplaceAll(line, []byte(""))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match(cleanString4, line); ok {
			reg := regexp.MustCompile(cleanString4)
			newByte := reg.ReplaceAll(line, []byte(""))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
}

func removeHeader(path, newPath string, lineNumber int) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	info, _ := os.Stat(path)
	mode := info.Mode()
	array := strings.Split(string(file), "\n")
	array = array[lineNumber:]
	ioutil.WriteFile(newPath, []byte(strings.Join(array, "\n")), mode)
	//fmt.Println("remove line successful")
}
