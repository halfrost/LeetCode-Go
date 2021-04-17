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

func isSymmetric(root *TreeNode) bool {
	var dfs func(rootLeft, rootRight *TreeNode) bool
	dfs = func(rootLeft, rootRight *TreeNode) bool {
		if rootLeft == nil && rootRight == nil {
			return true
		}
		if rootLeft == nil || rootRight == nil {
			return false
		}
		if rootLeft.Val != rootRight.Val {
			return false
		}
		return dfs(rootLeft.Left, rootRight.Right) && dfs(rootLeft.Right, rootRight.Left)
	}
	return root == nil || dfs(root.Left, root.Right)
}
