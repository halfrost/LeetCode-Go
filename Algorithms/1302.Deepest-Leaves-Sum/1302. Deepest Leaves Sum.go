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
func deepestLeavesSum(root *TreeNode) int {
	maxLevel, sum := 0, 0
	dfsDeepestLeavesSum(root, 0, &maxLevel, &sum)
	return sum
}

func dfsDeepestLeavesSum(root *TreeNode, level int, maxLevel, sum *int) {
	if root == nil {
		return
	}
	if level > *maxLevel {
		*maxLevel, *sum = level, root.Val
	} else if level == *maxLevel {
		*sum += root.Val
	}
	dfsDeepestLeavesSum(root.Left, level+1, maxLevel, sum)
	dfsDeepestLeavesSum(root.Right, level+1, maxLevel, sum)
}
