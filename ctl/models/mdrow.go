package models

import (
	"fmt"
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
