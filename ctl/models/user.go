package models

import (
	"fmt"
)

// UserInfo define
type UserInfo struct {
	UserName         string `json:"user_name"`
	NumSolved        int32  `json:"num_solved"`
	NumTotal         int32  `json:"num_total"`
	AcEasy           int32  `json:"ac_easy"`
	AcMedium         int32  `json:"ac_medium"`
	AcHard           int32  `json:"ac_hard"`
	EasyTotal        int32
	MediumTotal      int32
	HardTotal        int32
	OptimizingEasy   int32
	OptimizingMedium int32
	OptimizingHard   int32
	FrequencyHigh    float64 `json:"frequency_high"`
	FrequencyMid     float64 `json:"frequency_mid"`
	CategorySlug     string  `json:"category_slug"`
}

// |    |  Easy  |  Medium  |  Hard |  Total |  optimizing |
// |:--------:|:--------------------------------------------------------------|:--------:|:--------:|:--------:|:--------:|
func (ui UserInfo) table() string {
	res := "|    |  Easy  |  Medium  |  Hard |  Total |\n"
	res += "|:--------:|:--------:|:--------:|:--------:|:--------:|\n"
	res += fmt.Sprintf("|Optimizing|%v|%v|%v|%v|\n", ui.OptimizingEasy, ui.OptimizingMedium, ui.OptimizingHard, ui.OptimizingEasy+ui.OptimizingMedium+ui.OptimizingHard)
	res += fmt.Sprintf("|Accepted|**%v**|**%v**|**%v**|**%v**|\n", ui.AcEasy, ui.AcMedium, ui.AcHard, ui.AcEasy+ui.AcMedium+ui.AcHard)
	res += fmt.Sprintf("|Total|%v|%v|%v|%v|\n", ui.EasyTotal, ui.MediumTotal, ui.HardTotal, ui.EasyTotal+ui.MediumTotal+ui.HardTotal)
	res += fmt.Sprintf("|Perfection Rate|%.1f%%|%.1f%%|%.1f%%|%.1f%%|\n", (1-float64(ui.OptimizingEasy)/float64(ui.AcEasy))*100, (1-float64(ui.OptimizingMedium)/float64(ui.AcMedium))*100, (1-float64(ui.OptimizingHard)/float64(ui.AcHard))*100, (1-float64(ui.OptimizingEasy+ui.OptimizingMedium+ui.OptimizingHard)/float64(ui.AcEasy+ui.AcMedium+ui.AcHard))*100)
	res += fmt.Sprintf("|Completion Rate|%.1f%%|%.1f%%|%.1f%%|%.1f%%|\n", float64(ui.AcEasy)/float64(ui.EasyTotal)*100, float64(ui.AcMedium)/float64(ui.MediumTotal)*100, float64(ui.AcHard)/float64(ui.HardTotal)*100, float64(ui.AcEasy+ui.AcMedium+ui.AcHard)/float64(ui.EasyTotal+ui.MediumTotal+ui.HardTotal)*100)
	// 加这一行是为了撑开整个表格
	res += "|------------|----------------------------|----------------------------|----------------------------|----------------------------|"
	return res
}

// PersonalData define
func (ui UserInfo) PersonalData() string {
	return ui.table()
}
