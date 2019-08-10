package leetcode

import (
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

// 解法二 模拟，时间复杂度 O(n)
func maximumProduct1(nums []int) int {
	n1, n2, n3 := -1<<63, -1<<63, -1<<63
	n4, n5 := 0, 0
	for _, v := range nums {
		if v > n1 {
			n3 = n2
			n2 = n1
			n1 = v
		} else if v > n2 {
			n3 = n2
			n2 = v
		} else if v > n3 {
			n3 = v
		}
		if v < n4 {
			n5 = n4
			n4 = v
		} else if v < n5 {
			n5 = v
		}
	}
	if n2*n3 > n4*n5 {
		return n1 * n2 * n3
	}
	return n1 * n4 * n5
}
