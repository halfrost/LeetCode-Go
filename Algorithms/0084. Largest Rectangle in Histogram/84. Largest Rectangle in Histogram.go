package leetcode

import "fmt"

func largestRectangleArea(heights []int) int {
	maxArea, stack, height := 0, []int{}, 0
	for i := 0; i <= len(heights); i++ {
		if i == len(heights) {
			height = 0
		} else {
			height = heights[i]
		}
		if len(stack) == 0 || height >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			tmp := stack[len(stack)-1]
			fmt.Printf("1. tmp = %v stack = %v\n", tmp, stack)
			stack = stack[:len(stack)-1]
			length := 0
			if len(stack) == 0 {
				length = i
			} else {
				length = i - 1 - stack[len(stack)-1]
				fmt.Printf("2. length = %v stack = %v i = %v\n", length, stack, i)
			}
			maxArea = max(maxArea, heights[tmp]*length)
			fmt.Printf("3. maxArea = %v heights[tmp]*length = %v\n", maxArea, heights[tmp]*length)
			i--
		}
	}
	return maxArea
}
