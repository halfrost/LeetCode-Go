package models

type LeetCodeProblemAll struct {
	UserName        string            `json:"user_name"`
	NumSolved       int32             `json:"num_solved"`
	NumTotal        int32             `json:"num_total"`
	AcEasy          int32             `json:"ac_easy"`
	AcMedium        int32             `json:"ac_medium"`
	AcHard          int32             `json:"ac_hard"`
	StatStatusPairs []StatStatusPairs `json:"stat_status_pairs"`
	FrequencyHigh   int32             `json:"frequency_high"`
	FrequencyMid    int32             `json:"frequency_mid"`
	CategorySlug    string            `json:"category_slug"`
}

type StatStatusPairs struct {
	Stat       Stat       `json:"stat"`
	Difficulty Difficulty `json:"difficulty"`
	PaidOnly   bool       `json:"paid_only"`
	IsFavor    bool       `json:"is_favor"`
	Frequency  float64    `json:"frequency"`
	Progress   float64    `json:"progress"`
}

type Stat struct {
	QuestionTitle      string  `json:"question__title"`
	QuestionTitleSlug  string  `json:"question__title_slug"`
	TotalAcs           float64 `json:"total_acs"`
	TotalSubmitted     float64 `json:"total_submitted"`
	Acceptance         string
	Difficulty         string
	FrontendQuestionId int32 `json:"frontend_question_id"`
}

type Difficulty struct {
	Level int32 `json:"level"`
}

var DifficultyMap = map[int32]string{
	1: "Easy",
	2: "Medium",
	3: "Hard",
}
