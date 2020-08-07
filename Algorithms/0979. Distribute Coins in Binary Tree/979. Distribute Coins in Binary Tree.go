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

func distributeCoins(root *TreeNode) int {
	res := 0
	distributeCoinsDFS(root, &res)
	return res
}

func distributeCoinsDFS(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left, right := distributeCoinsDFS(root.Left, res), distributeCoinsDFS(root.Right, res)
	*res += abs(left) + abs(right)
	return left + right + root.Val - 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
