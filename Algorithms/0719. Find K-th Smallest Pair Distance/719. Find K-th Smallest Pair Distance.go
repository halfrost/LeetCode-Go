package leetcode

import (
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	low, high := 0, nums[len(nums)-1]-nums[0]
	for low < high {
		mid := low + (high-low)>>1
		tmp := findDistanceCount(nums, mid)
		if tmp >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// 解法一 双指针
func findDistanceCount(nums []int, num int) int {
	count, i := 0, 0
	for j := 1; j < len(nums); j++ {
		for nums[j]-nums[i] > num && i < j {
			i++
		}
		count += (j - i)
	}
	return count
}

// 解法二 暴力查找
func findDistanceCount1(nums []int, num int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] <= num {
				count++
			}
		}
	}
	return count
}
