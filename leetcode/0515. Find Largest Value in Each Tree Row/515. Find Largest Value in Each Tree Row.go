package leetcode

import (
	"math"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法一 层序遍历二叉树，再将每层排序取出最大值
func largestValues(root *TreeNode) []int {
	tmp := levelOrder(root)
	res := []int{}
	for i := 0; i < len(tmp); i++ {
		sort.Ints(tmp[i])
		res = append(res, tmp[i][len(tmp[i])-1])
	}
	return res
}

// 解法二 层序遍历二叉树，遍历过程中不断更新最大值
func largestValues1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := []*TreeNode{root}
	var res []int
	for len(q) > 0 {
		qlen := len(q)
		max := math.MinInt32
		for i := 0; i < qlen; i++ {
			node := q[0]
			q = q[1:]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, max)
	}
	return res
}
