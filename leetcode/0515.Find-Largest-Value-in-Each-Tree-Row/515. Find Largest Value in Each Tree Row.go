package leetcode

import (
	"math"
	"sort"
)

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

// this is 102 solution
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum, res, tmp := 1, 0, [][]int{}, []int{}
	for len(queue) != 0 {
		if curNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		if curNum == 0 {
			res = append(res, tmp)
			curNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}
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
