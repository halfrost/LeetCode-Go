package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

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
func sumRootToLeaf(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}
		sum = sum<<1 | node.Val
		// 上一行也可以写作 sum = sum*2 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return dfs(node.Left, sum) + dfs(node.Right, sum)
	}
	return dfs(root, 0)
}
