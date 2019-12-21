package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	findTiltDFS(root, &sum)
	return sum
}

func findTiltDFS(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	left := findTiltDFS(root.Left, sum)
	right := findTiltDFS(root.Right, sum)
	*sum += int(math.Abs(float64(left) - float64(right)))
	return root.Val + left + right
}
