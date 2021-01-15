package main

import (
	m "github.com/halfrost/LeetCode-Go/ctl/models"
	"github.com/halfrost/LeetCode-Go/ctl/util"
	"sort"
)

func statisticalData(problemsMap map[int]m.StatStatusPairs, solutionIds []int) (easyTotal, mediumTotal, hardTotal, optimizingEasy, optimizingMedium, optimizingHard int32, optimizingIds []int) {
	easyTotal, mediumTotal, hardTotal, optimizingEasy, optimizingMedium, optimizingHard, optimizingIds = 0, 0, 0, 0, 0, 0, []int{}
	for _, v := range problemsMap {
		switch m.DifficultyMap[v.Difficulty.Level] {
		case "Easy":
			{
				easyTotal++
				if v.Status == "ac" && util.BinarySearch(solutionIds, int(v.Stat.FrontendQuestionID)) == -1 {
					optimizingEasy++
					optimizingIds = append(optimizingIds, int(v.Stat.FrontendQuestionID))
				}
			}
		case "Medium":
			{
				mediumTotal++
				if v.Status == "ac" && util.BinarySearch(solutionIds, int(v.Stat.FrontendQuestionID)) == -1 {
					optimizingMedium++
					optimizingIds = append(optimizingIds, int(v.Stat.FrontendQuestionID))
				}
			}
		case "Hard":
			{
				hardTotal++
				if v.Status == "ac" && util.BinarySearch(solutionIds, int(v.Stat.FrontendQuestionID)) == -1 {
					optimizingHard++
					optimizingIds = append(optimizingIds, int(v.Stat.FrontendQuestionID))
				}
			}
		}
	}
	sort.Ints(optimizingIds)
	return easyTotal, mediumTotal, hardTotal, optimizingEasy, optimizingMedium, optimizingHard, optimizingIds
}
