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

// 解法一
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum, res, tmp, curDir := 1, 0, [][]int{}, []int{}, 0
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
			if curDir == 1 {
				for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
					tmp[i], tmp[j] = tmp[j], tmp[i]
				}
			}
			res = append(res, tmp)
			curNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
			if curDir == 0 {
				curDir = 1
			} else {
				curDir = 0
			}
		}
	}
	return res
}

// 解法二 递归
func zigzagLevelOrder0(root *TreeNode) [][]int {
	var res [][]int
	search(root, 0, &res)
	return res
}

func search(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}
	for len(*res) < depth+1 {
		*res = append(*res, []int{})
	}
	if depth%2 == 0 {
		(*res)[depth] = append((*res)[depth], root.Val)
	} else {
		(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
	}
	search(root.Left, depth+1, res)
	search(root.Right, depth+1, res)
}

// 解法三 BFS
func zigzagLevelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	size, i, j, lay, tmp, flag := 0, 0, 0, []int{}, []*TreeNode{}, false
	for len(q) > 0 {
		size = len(q)
		tmp = []*TreeNode{}
		lay = make([]int, size)
		j = size - 1
		for i = 0; i < size; i++ {
			root = q[0]
			q = q[1:]
			if !flag {
				lay[i] = root.Val
			} else {
				lay[j] = root.Val
				j--
			}
			if root.Left != nil {
				tmp = append(tmp, root.Left)
			}
			if root.Right != nil {
				tmp = append(tmp, root.Right)
			}

		}
		res = append(res, lay)
		flag = !flag
		q = tmp
	}
	return res
}
