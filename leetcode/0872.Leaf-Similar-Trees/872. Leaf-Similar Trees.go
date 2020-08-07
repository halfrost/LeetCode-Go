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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1, leaf2 := []int{}, []int{}
	dfsLeaf(root1, &leaf1)
	dfsLeaf(root2, &leaf2)
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := range leaf1 {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}

func dfsLeaf(root *TreeNode, leaf *[]int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*leaf = append(*leaf, root.Val)
		}
		dfsLeaf(root.Left, leaf)
		dfsLeaf(root.Right, leaf)
	}
}
