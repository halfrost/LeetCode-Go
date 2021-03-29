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

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	res, index := make([]int, 0, len(voyage)), 0
	if travelTree(root, &index, voyage, &res) {
		return res
	}
	return []int{-1}
}

func travelTree(root *TreeNode, index *int, voyage []int, res *[]int) bool {
	if root == nil {
		return true
	}
	if root.Val != voyage[*index] {
		return false
	}
	*index++
	if root.Left != nil && root.Left.Val != voyage[*index] {
		*res = append(*res, root.Val)
		return travelTree(root.Right, index, voyage, res) && travelTree(root.Left, index, voyage, res)
	}
	return travelTree(root.Left, index, voyage, res) && travelTree(root.Right, index, voyage, res)
}
