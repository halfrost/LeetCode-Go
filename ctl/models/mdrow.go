package models

import (
	"fmt"
	"strings"
)

// Mdrow define
type Mdrow struct {
	FrontendQuestionID int32  `json:"question_id"`
	QuestionTitle      string `json:"question__title"`
	QuestionTitleSlug  string `json:"question__title_slug"`
	SolutionPath       string `json:"solution_path"`
	Acceptance         string `json:"acceptance"`
	Difficulty         string `json:"difficulty"`
	Frequency          string `json:"frequency"`
}

// GenerateMdRows define
func GenerateMdRows(solutionIds []int, mdrows []Mdrow) {
	for i := 0; i < len(solutionIds); i++ {
		id := mdrows[solutionIds[i]-1].FrontendQuestionID
		if solutionIds[i] == int(id) {
			//fmt.Printf("id = %v i = %v solutionIds = %v\n", id, i, solutionIds[i])
			mdrows[id-1].SolutionPath = fmt.Sprintf("[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/%v)", fmt.Sprintf("%04d.%v", id, strings.Replace(strings.TrimSpace(mdrows[id-1].QuestionTitle), " ", "-", -1)))
		} else {
			fmt.Printf("序号出错了 solutionIds = %v id = %v\n", solutionIds[i], id)
		}
	}
}

// | 0001 | Two Sum  | [Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0001.Two-Sum)| 45.6%  | Easy | |
func (m Mdrow) tableLine() string {
	return fmt.Sprintf("|%04d|%v|%v|%v|%v||\n", m.FrontendQuestionID, m.QuestionTitle, m.SolutionPath, m.Acceptance, m.Difficulty)
}

// SortByQuestionID define
type SortByQuestionID []Mdrow

func (a SortByQuestionID) Len() int      { return len(a) }
func (a SortByQuestionID) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByQuestionID) Less(i, j int) bool {
	return a[i].FrontendQuestionID < a[j].FrontendQuestionID
}

// Mdrows define
type Mdrows struct {
	Mdrows []Mdrow
}

// | No.    |  Title  |  Solution  |  Acceptance |  Difficulty |  Frequency |
// |:--------:|:--------------------------------------------------------------|:--------:|:--------:|:--------:|:--------:|
func (mds Mdrows) table() string {
	res := "| No.    |  Title  |  Solution  |  Acceptance |  Difficulty |  Frequency |\n"
	res += "|:--------:|:--------------------------------------------------------------|:--------:|:--------:|:--------:|:--------:|\n"
	for _, p := range mds.Mdrows {
		res += p.tableLine()
	}
	// 加这一行是为了撑开整个表格
	res += "|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|"
	return res
}

// AvailableTable define
func (mds Mdrows) AvailableTable() string {
	return mds.table()
}
