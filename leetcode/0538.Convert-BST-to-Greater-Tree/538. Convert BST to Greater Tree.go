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

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	dfs538(root, &sum)
	return root
}

func dfs538(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	dfs538(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	dfs538(root.Left, sum)
}
