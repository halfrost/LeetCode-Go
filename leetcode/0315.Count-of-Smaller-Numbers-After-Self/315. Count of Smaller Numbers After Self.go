package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/template"
)

// 解法一 线段树
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

// 解法二 树状数组
func countSmaller1(nums []int) []int {
	// copy 一份原数组至所有数字 allNums 数组中
	allNums, res := make([]int, len(nums)), []int{}
	copy(allNums, nums)
	// 将 allNums 离散化
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i < len(allNums); i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}
	// 树状数组 Query
	bit := template.BinaryIndexedTree{}
	bit.Init(k)
	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, bit.Query(kth[nums[i]]-1))
		bit.Add(kth[nums[i]], 1)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
