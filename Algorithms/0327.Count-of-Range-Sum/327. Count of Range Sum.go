package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/template"
)

// 解法一 线段树，时间复杂度 O(n log n)
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}
	st, prefixSum, sumMap, sumArray, res := template.SegmentCountTree{}, make([]int, len(nums)), make(map[int]int, 0), []int{}, 0
	prefixSum[0], sumMap[nums[0]] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
		sumMap[prefixSum[i]] = prefixSum[i]
	}
	// sumArray 是 prefixSum 去重之后的版本，利用 sumMap 去重
	for _, v := range sumMap {
		sumArray = append(sumArray, v)
	}
	// 排序是为了使得线段树中的区间 left <= right，如果此处不排序，线段树中的区间有很多不合法。
	sort.Ints(sumArray)
	// 初始化线段树，节点内的值都赋值为 0，即计数为 0
	st.Init(sumArray, func(i, j int) int {
		return 0
	})
	// 倒序是为了方便寻找 j，sum(i，j) 规定了 j >= i，所以倒序遍历，i 从大到小
	for i := len(nums) - 1; i >= 0; i-- {
		// 插入的 prefixSum[i] 即是 j
		st.UpdateCount(prefixSum[i])
		if i > 0 {
			res += st.Query(lower+prefixSum[i-1], upper+prefixSum[i-1])
		} else {
			res += st.Query(lower, upper)
		}
	}
	return res
}

// 解法二 暴力，时间复杂度 O(n^2)
func countRangeSum1(nums []int, lower int, upper int) int {
	res, n := 0, len(nums)
	for i := 0; i < n; i++ {
		tmp := 0
		for j := i; j < n; j++ {
			if i == j {
				tmp = nums[i]
			} else {
				tmp += nums[j]
			}
			if tmp <= upper && tmp >= lower {
				res++
			}
		}
	}
	return res
}
