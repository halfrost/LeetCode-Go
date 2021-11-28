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

func isCompleteTree(root *TreeNode) bool {
	queue, found := []*TreeNode{root}, false
	for len(queue) > 0 {
		node := queue[0] //取出每一层的第一个节点
		queue = queue[1:]
		if node == nil {
			found = true
		} else {
			if found {
				return false // 层序遍历中，两个不为空的节点中出现一个 nil
			}
			//如果左孩子为nil，则append进去的node.Left为nil
			queue = append(queue, node.Left, node.Right)
		}
	}
	return true
}
