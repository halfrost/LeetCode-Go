package leetcode

// 解法一 二分
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}
		mid := low + (high-low)>>1
		if nums[mid] >= nums[low] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

// 解法二 二分
func findMin1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[low] < nums[high] {
			return nums[low]
		}
		if (mid == len(nums)-1 && nums[mid-1] > nums[mid]) || (mid < len(nums)-1 && mid > 0 && nums[mid-1] > nums[mid] && nums[mid] < nums[mid+1]) {
			return nums[mid]
		}
		if nums[mid] > nums[low] && nums[low] > nums[high] { // mid 在数值大的一部分区间里
			low = mid + 1
		} else if nums[mid] < nums[low] && nums[low] > nums[high] { // mid 在数值小的一部分区间里
			high = mid - 1
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}

// 解法三 暴力
func findMin2(nums []int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if min > num {
			min = num
		}
	}
	return min
}
