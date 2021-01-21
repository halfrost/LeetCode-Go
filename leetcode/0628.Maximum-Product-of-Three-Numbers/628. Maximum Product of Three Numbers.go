package leetcode

import (
	"math"
	"sort"
)

// 解法一 排序，时间复杂度 O(n log n)
func maximumProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	if len(nums) <= 3 {
		for i := 0; i < len(nums); i++ {
			res = res * nums[i]
		}
		return res
	}
	sort.Ints(nums)
	if nums[len(nums)-1] <= 0 {
		return 0
	}
	return max(nums[0]*nums[1]*nums[len(nums)-1], nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法二 模拟，时间复杂度 O(n)
func maximumProduct1(nums []int) int {
	max := make([]int, 0)
	max = append(max, math.MinInt64, math.MinInt64, math.MinInt64)
	min := make([]int, 0)
	min = append(min, math.MaxInt64, math.MaxInt64)
	for _, num := range nums {
		if num > max[0] {
			max[0], max[1], max[2] = num, max[0], max[1]
		} else if num > max[1] {
			max[1], max[2] = num, max[1]
		} else if num > max[2] {
			max[2] = num
		}
		if num < min[0] {
			min[0], min[1] = num, min[0]
		} else if num < min[1] {
			min[1] = num
		}
	}
	maxProduct1, maxProduct2 := min[0]*min[1]*max[0], max[0]*max[1]*max[2]
	if maxProduct1 > maxProduct2 {
		return maxProduct1
	}
	return maxProduct2
}
