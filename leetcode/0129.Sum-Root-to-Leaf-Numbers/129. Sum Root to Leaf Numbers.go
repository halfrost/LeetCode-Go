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

func sumNumbers(root *TreeNode) int {
	res := 0
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += sum
		return
	}
	dfs(root.Left, sum, res)
	dfs(root.Right, sum, res)
}
