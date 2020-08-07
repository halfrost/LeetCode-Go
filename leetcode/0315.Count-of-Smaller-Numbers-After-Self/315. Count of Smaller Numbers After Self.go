package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/template"
)

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	st, minNum, numsMap, numsArray, res := template.SegmentCountTree{}, 0, make(map[int]int, 0), []int{}, make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = nums[i]
	}
	for _, v := range numsMap {
		numsArray = append(numsArray, v)
	}
	// 排序是为了使得线段树中的区间 left <= right，如果此处不排序，线段树中的区间有很多不合法。
	sort.Ints(numsArray)
	minNum = numsArray[0]
	// 初始化线段树，节点内的值都赋值为 0，即计数为 0
	st.Init(numsArray, func(i, j int) int {
		return 0
	})
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == minNum {
			res[i] = 0
			st.UpdateCount(nums[i])
			continue
		}
		st.UpdateCount(nums[i])
		res[i] = st.Query(minNum, nums[i]-1)
	}
	return res
}
