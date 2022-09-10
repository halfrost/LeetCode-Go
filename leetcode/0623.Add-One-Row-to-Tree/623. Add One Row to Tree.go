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
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		tmp := &TreeNode{Val: v, Left: root, Right: nil}
		return tmp
	}
	level := 1
	addTreeRow(root, v, d, &level)
	return root
}

func addTreeRow(root *TreeNode, v, d int, currLevel *int) {
	if *currLevel == d-1 {
		root.Left = &TreeNode{Val: v, Left: root.Left, Right: nil}
		root.Right = &TreeNode{Val: v, Left: nil, Right: root.Right}
		return
	}
	*currLevel++
	if root.Left != nil {
		addTreeRow(root.Left, v, d, currLevel)
	}
	if root.Right != nil {
		addTreeRow(root.Right, v, d, currLevel)
	}
	*currLevel--
}
