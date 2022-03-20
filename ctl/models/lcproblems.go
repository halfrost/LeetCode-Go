package models

import (
	"fmt"
	"strings"
)

// LeetCodeProblemAll define
type LeetCodeProblemAll struct {
	UserName        string            `json:"user_name"`
	NumSolved       int32             `json:"num_solved"`
	NumTotal        int32             `json:"num_total"`
	AcEasy          int32             `json:"ac_easy"`
	AcMedium        int32             `json:"ac_medium"`
	AcHard          int32             `json:"ac_hard"`
	StatStatusPairs []StatStatusPairs `json:"stat_status_pairs"`
	FrequencyHigh   float64           `json:"frequency_high"`
	FrequencyMid    float64           `json:"frequency_mid"`
	CategorySlug    string            `json:"category_slug"`
	AcEasyTotal     int32
	AcMediumTotal   int32
	AcHardTotal     int32
}

// ConvertUserInfoModel define
func ConvertUserInfoModel(lpa LeetCodeProblemAll) UserInfo {
	info := UserInfo{}
	info.UserName = lpa.UserName
	info.NumSolved = lpa.NumSolved
	info.NumTotal = lpa.NumTotal
	info.AcEasy = lpa.AcEasy
	info.AcMedium = lpa.AcMedium
	info.AcHard = lpa.AcHard
	info.FrequencyHigh = lpa.FrequencyHigh
	info.FrequencyMid = lpa.FrequencyMid
	info.CategorySlug = lpa.CategorySlug
	return info
}

// StatStatusPairs define
type StatStatusPairs struct {
	Stat       Stat       `json:"stat"`
	Status     string     `json:"status"`
	Difficulty Difficulty `json:"difficulty"`
	PaidOnly   bool       `json:"paid_only"`
	IsFavor    bool       `json:"is_favor"`
	Frequency  float64    `json:"frequency"`
	Progress   float64    `json:"progress"`
}

// ConvertMdModelFromSsp define
func ConvertMdModelFromSsp(problems []StatStatusPairs) []Mdrow {
	mdrows := []Mdrow{}
	for _, problem := range problems {
		res := Mdrow{}
		res.FrontendQuestionID = problem.Stat.FrontendQuestionID
		res.QuestionTitle = strings.TrimSpace(problem.Stat.QuestionTitle)
		res.QuestionTitleSlug = strings.TrimSpace(problem.Stat.QuestionTitleSlug)
		res.Acceptance = fmt.Sprintf("%.1f%%", (problem.Stat.TotalAcs/problem.Stat.TotalSubmitted)*100)
		res.Difficulty = DifficultyMap[problem.Difficulty.Level]
		res.Frequency = fmt.Sprintf("%f", problem.Frequency)
		mdrows = append(mdrows, res)
	}
	return mdrows
}

// ConvertMdModelFromIds define
func ConvertMdModelFromIds(problemsMap map[int]StatStatusPairs, ids []int) []Mdrow {
	mdrows := []Mdrow{}
	for _, v := range ids {
		res, problem := Mdrow{}, problemsMap[v]
		res.FrontendQuestionID = problem.Stat.FrontendQuestionID
		res.QuestionTitle = strings.TrimSpace(problem.Stat.QuestionTitle)
		res.QuestionTitleSlug = strings.TrimSpace(problem.Stat.QuestionTitleSlug)
		res.Acceptance = fmt.Sprintf("%.1f%%", (problem.Stat.TotalAcs/problem.Stat.TotalSubmitted)*100)
		res.Difficulty = DifficultyMap[problem.Difficulty.Level]
		res.Frequency = fmt.Sprintf("%f", problem.Frequency)
		mdrows = append(mdrows, res)
	}
	return mdrows
}

// Stat define
type Stat struct {
	QuestionTitle      string  `json:"question__title"`
	QuestionTitleSlug  string  `json:"question__title_slug"`
	TotalAcs           float64 `json:"total_acs"`
	TotalSubmitted     float64 `json:"total_submitted"`
	Acceptance         string
	Difficulty         string
	FrontendQuestionID int32 `json:"frontend_question_id"`
}

// Difficulty define
type Difficulty struct {
	Level int32 `json:"level"`
}

// DifficultyMap define
var DifficultyMap = map[int32]string{
	1: "Easy",
	2: "Medium",
	3: "Hard",
}
