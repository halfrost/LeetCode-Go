package leetcode

import (
	"fmt"
	"math"
)

// 解法一 单调栈
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	num3, stack := math.MinInt64, []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < num3 {
			return true
		}
		for len(stack) != 0 && nums[i] > stack[len(stack)-1] {
			num3 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
		fmt.Printf("stack = %v \n", stack)
	}
	return false
}

// 解法二 暴力解法，超时！
func find132pattern1(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	for j := 0; j < len(nums); j++ {
		stack := []int{}
		for i := j; i < len(nums); i++ {
			if len(stack) == 0 || (len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]]) {
				stack = append(stack, i)
			} else if nums[i] < nums[stack[len(stack)-1]] {
				index := len(stack) - 1
				for ; index >= 0; index-- {
					if nums[stack[index]] < nums[i] {
						return true
					}
				}
			}
		}
	}
	return false
}
