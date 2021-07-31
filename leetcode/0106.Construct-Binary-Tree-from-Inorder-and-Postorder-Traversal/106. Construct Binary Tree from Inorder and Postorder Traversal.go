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

// 解法一, 直接传入需要的 slice 范围作为输入, 可以避免申请对应 inorder 索引的内存, 内存使用(leetcode test case) 4.7MB -> 4.3MB.
func buildTree(inorder []int, postorder []int) *TreeNode {
	postorderLen := len(postorder)
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[postorderLen-1]}
	postorder = postorder[:postorderLen-1]
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(inorder[:pos], postorder[:len(inorder[:pos])])
			root.Right = buildTree(inorder[pos+1:], postorder[len(inorder[:pos]):])
		}
	}
	return root
}

// 解法二
func buildTree1(inorder []int, postorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return buildInPos2TreeDFS(postorder, 0, len(postorder)-1, 0, inPos)
}

func buildInPos2TreeDFS(post []int, postStart int, postEnd int, inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	root := &TreeNode{Val: post[postEnd]}
	rootIdx := inPos[post[postEnd]]
	leftLen := rootIdx - inStart
	root.Left = buildInPos2TreeDFS(post, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = buildInPos2TreeDFS(post, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
	return root
}
