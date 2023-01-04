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
	mdMap := map[int]Mdrow{}
	for _, row := range mdrows {
		mdMap[int(row.FrontendQuestionID)] = row
	}
	for i := 0; i < len(solutionIds); i++ {
		if row, ok := mdMap[solutionIds[i]]; ok {
			s7 := standardizedTitle(row.QuestionTitle, row.FrontendQuestionID)
			mdMap[solutionIds[i]] = Mdrow{
				FrontendQuestionID: row.FrontendQuestionID,
				QuestionTitle:      strings.TrimSpace(row.QuestionTitle),
				QuestionTitleSlug:  row.QuestionTitleSlug,
				SolutionPath:       fmt.Sprintf("[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/%v)", fmt.Sprintf("%04d.%v", solutionIds[i], s7)),
				Acceptance:         row.Acceptance,
				Difficulty:         row.Difficulty,
				Frequency:          row.Frequency,
			}
		} else {
			fmt.Printf("序号不存在 len(solutionIds) = %v len(mdrows) = %v len(solutionIds) = %v solutionIds[i] = %v QuestionTitle = %v\n", len(solutionIds), len(mdrows), len(solutionIds), solutionIds[i], mdrows[solutionIds[i]-1].QuestionTitle)
		}
	}
	for i := range mdrows {
		mdrows[i] = Mdrow{
			FrontendQuestionID: mdrows[i].FrontendQuestionID,
			QuestionTitle:      strings.TrimSpace(mdrows[i].QuestionTitle),
			QuestionTitleSlug:  mdrows[i].QuestionTitleSlug,
			SolutionPath:       mdMap[int(mdrows[i].FrontendQuestionID)].SolutionPath,
			Acceptance:         mdrows[i].Acceptance,
			Difficulty:         mdrows[i].Difficulty,
			Frequency:          mdrows[i].Frequency,
		}
	}
	// fmt.Printf("mdrows = %v\n\n", mdrows)
}

// | 0001 | Two Sum  | [Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0001.Two-Sum)| 45.6%  | Easy | |
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
