package models

import (
	"strconv"
)

type Mdrow struct {
	FrontendQuestionId string `json:"question_id"`
	QuestionTitle      string `json:"question__title"`
	QuestionTitleSlug  string `json:"question__title_slug"`
	SolutionPath       string `json:"solution_path"`
	Acceptance         string `json:"acceptance"`
	Difficulty         string `json:"difficulty"`
	Frequency          string `json:"frequency"`
}

// SortByQuestionId define
type SortByQuestionId []Mdrow

func (a SortByQuestionId) Len() int      { return len(a) }
func (a SortByQuestionId) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByQuestionId) Less(i, j int) bool {
	first, _ := strconv.Atoi(a[i].FrontendQuestionId)
	second, _ := strconv.Atoi(a[j].FrontendQuestionId)
	return first < second
}
