package leetcode

import "sort"

// 解法一 快慢指针
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	walker := 0
	for walker != slow {
		walker = nums[walker]
		slow = nums[slow]
	}
	return walker
}

// 解法二 二分搜索
func findDuplicate1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// 解法三
func findDuplicate2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	diff := -1
	for i := 0; i < len(nums); i++ {
		if nums[i]-i-1 >= diff {
			diff = nums[i] - i - 1
		} else {
			return nums[i]
		}
	}
	return 0
}
