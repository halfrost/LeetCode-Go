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
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	visit := []int{}
	findDistanceK(root, target, K, &visit)
	return visit
}

func findDistanceK(root, target *TreeNode, K int, visit *[]int) int {
	if root == nil {
		return -1
	}
	if root == target {
		findChild(root, K, visit)
		return K - 1
	}
	leftDistance := findDistanceK(root.Left, target, K, visit)
	if leftDistance == 0 {
		findChild(root, leftDistance, visit)
	}
	if leftDistance > 0 {
		findChild(root.Right, leftDistance-1, visit)
		return leftDistance - 1
	}
	rightDistance := findDistanceK(root.Right, target, K, visit)
	if rightDistance == 0 {
		findChild(root, rightDistance, visit)
	}
	if rightDistance > 0 {
		findChild(root.Left, rightDistance-1, visit)
		return rightDistance - 1
	}
	return -1
}

func findChild(root *TreeNode, K int, visit *[]int) {
	if root == nil {
		return
	}
	if K == 0 {
		*visit = append(*visit, root.Val)
	} else {
		findChild(root.Left, K-1, visit)
		findChild(root.Right, K-1, visit)
	}
}
