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

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var left, right int
	dfsBtreeGameWinningMove(root, &left, &right, x)
	up := n - left - right - 1
	n /= 2
	return left > n || right > n || up > n
}

func dfsBtreeGameWinningMove(node *TreeNode, left, right *int, x int) int {
	if node == nil {
		return 0
	}
	l, r := dfsBtreeGameWinningMove(node.Left, left, right, x), dfsBtreeGameWinningMove(node.Right, left, right, x)
	if node.Val == x {
		*left, *right = l, r
	}
	return l + r + 1
}
