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
func rob337(root *TreeNode) int {
	a, b := dfsTreeRob(root)
	return max(a, b)
}

func dfsTreeRob(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}
	l0, l1 := dfsTreeRob(root.Left)
	r0, r1 := dfsTreeRob(root.Right)
	// 当前节点没有被打劫
	tmp0 := max(l0, l1) + max(r0, r1)
	// 当前节点被打劫
	tmp1 := root.Val + l0 + r0
	return tmp0, tmp1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
