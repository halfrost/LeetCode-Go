package leetcode

import "sort"

// 解法一 排序，排序的方法反而速度是最快的
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 解法二 这个方法的理论依据是 partition 得到的点的下标就是最终排序之后的下标，根据这个下标，我们可以判断第 K 大的数在哪里
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	return selection(nums, 0, len(nums)-1, len(nums)-k)
}

func selection(arr []int, l, r, k int) int {
	if l == r {
		return arr[l]
	}
	p := partition164(arr, l, r)
	if k == p {
		return arr[p]
	} else if k < p {
		return selection(arr, l, p-1, k)
	} else {
		return selection(arr, p+1, r, k)
	}
}
