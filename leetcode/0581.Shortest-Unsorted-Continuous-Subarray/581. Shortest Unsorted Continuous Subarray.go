package leetcode

import "math"

func findUnsortedSubarray(nums []int) int {
	n, left, right, minR, maxL, isSort := len(nums), -1, -1, math.MaxInt32, math.MinInt32, false
	// left
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			isSort = true
		}
		if isSort {
			minR = min(minR, nums[i])
		}
	}
	isSort = false
	// right
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			isSort = true
		}
		if isSort {
			maxL = max(maxL, nums[i])
		}
	}
	// minR
	for i := 0; i < n; i++ {
		if nums[i] > minR {
			left = i
			break
		}
	}
	// maxL
	for i := n - 1; i >= 0; i-- {
		if nums[i] < maxL {
			right = i
			break
		}
	}
	if left == -1 || right == -1 {
		return 0
	}
	return right - left + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
