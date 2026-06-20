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

// 解法一 DFS
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res, maxHeight := 0, -1
	findBottomLeftValueDFS(root, 0, &res, &maxHeight)
	return res
}

func findBottomLeftValueDFS(root *TreeNode, curHeight int, res, maxHeight *int) {
	if curHeight > *maxHeight && root.Left == nil && root.Right == nil {
		*maxHeight = curHeight
		*res = root.Val
	}
	if root.Left != nil {
		findBottomLeftValueDFS(root.Left, curHeight+1, res, maxHeight)
	}
	if root.Right != nil {
		findBottomLeftValueDFS(root.Right, curHeight+1, res, maxHeight)
	}
}

// 解法二 BFS
func findBottomLeftValue1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res, queue := root.Val, []*TreeNode{root}
	for len(queue) > 0 {
		res = queue[0].Val // 每层第一个节点即该层最左节点，最后停留的就是底层最左值
		next := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		queue = next
	}
	return res
}
