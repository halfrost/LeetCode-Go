package leetcode

import "math"

func maxSubarraySumCircular(nums []int) int {
	var max1, max2, sum int

	// case: no circulation
	max1 = int(math.Inf(-1))
	l := len(nums)
	for i := 0; i < l; i++ {
		sum += nums[i]
		if sum > max1 {
			max1 = sum
		}
		if sum < 1 {
			sum = 0
		}
	}

	// case: circling
	arr_sum := 0
	for i := 0; i < l; i++ {
		arr_sum += nums[i]
	}

	sum = 0
	min_sum := 0
	for i := 1; i < l-1; i++ {
		sum += nums[i]
		if sum >= 0 {
			sum = 0
		}
		if sum < min_sum {
			min_sum = sum
		}
	}
	max2 = arr_sum - min_sum

	return max(max1, max2)
}

func max(nums ...int) int {
	max := int(math.Inf(-1))
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
