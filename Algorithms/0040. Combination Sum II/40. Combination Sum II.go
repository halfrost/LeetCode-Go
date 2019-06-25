package leetcode

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates) // 这里是去重的关键逻辑
	findcombinationSum2(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum2(nums []int, target, index int, c []int, res *[][]int) {
	if target == 0 {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] { // 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
			continue
		}
		if target >= nums[i] {
			c = append(c, nums[i])
			findcombinationSum2(nums, target-nums[i], i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
