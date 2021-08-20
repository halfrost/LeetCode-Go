package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	checkDiameter(root, &result)
	return result
}

func checkDiameter(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := checkDiameter(root.Left, result)
	right := checkDiameter(root.Right, result)
	*result = max(*result, left+right)
	return max(left, right) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
