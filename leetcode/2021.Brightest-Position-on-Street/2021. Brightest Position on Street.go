package leetcode

import (
	"sort"
)

type lightItem struct {
	index int
	sign  int
}

func brightestPosition(lights [][]int) int {
	lightMap, lightItems := map[int]int{}, []lightItem{}
	for _, light := range lights {
		lightMap[light[0]-light[1]] += 1
		lightMap[light[0]+light[1]+1] -= 1
	}
	for k, v := range lightMap {
		lightItems = append(lightItems, lightItem{index: k, sign: v})
	}
	sort.SliceStable(lightItems, func(i, j int) bool {
		return lightItems[i].index < lightItems[j].index
	})
	res, border, tmp := 0, 0, 0
	for _, v := range lightItems {
		tmp += v.sign
		if border < tmp {
			res = v.index
			border = tmp
		}
	}
	return res
}
