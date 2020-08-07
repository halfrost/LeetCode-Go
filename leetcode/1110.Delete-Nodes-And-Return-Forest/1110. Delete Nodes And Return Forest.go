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
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	res, deleteMap := []*TreeNode{}, map[int]bool{}
	for _, v := range toDelete {
		deleteMap[v] = true
	}
	dfsDelNodes(root, deleteMap, true, &res)
	return res
}

func dfsDelNodes(root *TreeNode, toDel map[int]bool, isRoot bool, res *[]*TreeNode) bool {
	if root == nil {
		return false
	}
	if isRoot && !toDel[root.Val] {
		*res = append(*res, root)
	}
	isRoot = false
	if toDel[root.Val] {
		isRoot = true
	}
	if dfsDelNodes(root.Left, toDel, isRoot, res) {
		root.Left = nil
	}
	if dfsDelNodes(root.Right, toDel, isRoot, res) {
		root.Right = nil
	}
	return isRoot
}
