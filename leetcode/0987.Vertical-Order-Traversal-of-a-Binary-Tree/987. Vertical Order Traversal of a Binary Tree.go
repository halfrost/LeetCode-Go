package leetcode

import (
	"math"
	"sort"

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

type node struct {
	x, y, val int
}

func verticalTraversal(root *TreeNode) [][]int {
	var dfs func(root *TreeNode, x, y int)
	var nodes []node
	dfs = func(root *TreeNode, x, y int) {
		if root == nil {
			return
		}
		nodes = append(nodes, node{x, y, root.Val})
		dfs(root.Left, x+1, y-1)
		dfs(root.Right, x+1, y+1)
	}
	dfs(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]
		return a.y < b.y || a.y == b.y &&
			(a.x < b.x || a.x == b.x && a.val < b.val)
	})

	var res [][]int
	lastY := math.MinInt32
	for _, node := range nodes {
		if lastY != node.y {
			res = append(res, []int{node.val})
			lastY = node.y
		} else {
			res[len(res)-1] = append(res[len(res)-1], node.val)
		}
	}
	return res
}
