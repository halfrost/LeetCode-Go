package leetcode

// 解法一 逐行把矩阵看成柱状图，对每一行求“柱状图中最大矩形”（第 84 题）。
// 时间复杂度 O(m*n)，空间复杂度 O(n)。
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix[0])
	heights, res := make([]int, n), 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		if area := largestRectangleArea85(heights); area > res {
			res = area
		}
	}
	return res
}

// 单调栈求柱状图中的最大矩形面积。
func largestRectangleArea85(heights []int) int {
	res, stack := 0, []int{}
	for i := 0; i <= len(heights); i++ {
		cur := 0
		if i < len(heights) {
			cur = heights[i]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= cur {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			if area := heights[top] * width; area > res {
				res = area
			}
		}
		stack = append(stack, i)
	}
	return res
}
