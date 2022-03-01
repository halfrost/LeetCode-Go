package leetcode

import (
	"sort"
)

type node struct {
	value int
	count int
}

func minimumOperations(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	res, odd, even, oddMap, evenMap := 0, []node{}, []node{}, map[int]int{}, map[int]int{}

	for i := 0; i < len(nums); i += 2 {
		evenMap[nums[i]]++
	}
	for k, v := range evenMap {
		even = append(even, node{value: k, count: v})
	}
	sort.Slice(even, func(i, j int) bool {
		return even[i].count > even[j].count
	})

	for i := 1; i < len(nums); i += 2 {
		oddMap[nums[i]]++
	}
	for k, v := range oddMap {
		odd = append(odd, node{value: k, count: v})
	}
	sort.Slice(odd, func(i, j int) bool {
		return odd[i].count > odd[j].count
	})

	if even[0].value == odd[0].value {
		if len(even) == 1 && len(odd) != 1 {
			res = len(nums) - even[0].count - odd[1].count
		} else if len(odd) == 1 && len(even) != 1 {
			res = len(nums) - odd[0].count - even[1].count
		} else if len(odd) == 1 && len(even) == 1 {
			res = len(nums) / 2
		} else {
			// both != 1
			res = min(len(nums)-odd[0].count-even[1].count, len(nums)-odd[1].count-even[0].count)
		}
	} else {
		res = len(nums) - even[0].count - odd[0].count
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
