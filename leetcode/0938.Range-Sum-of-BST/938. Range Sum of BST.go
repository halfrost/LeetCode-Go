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

func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	preOrder(root, low, high, &res)
	return res
}

func preOrder(root *TreeNode, low, high int, res *int) {
	if root == nil {
		return
	}
	if low <= root.Val && root.Val <= high {
		*res += root.Val
	}
	preOrder(root.Left, low, high, res)
	preOrder(root.Right, low, high, res)
}
