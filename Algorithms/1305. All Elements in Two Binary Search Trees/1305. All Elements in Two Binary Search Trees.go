package leetcode

import (
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

// 解法一 合并排序
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := inorderTraversal(root1)
	arr2 := inorderTraversal(root2)
	arr1 = append(arr1, make([]int, len(arr2))...)
	merge(arr1, len(arr1)-len(arr2), arr2, len(arr2))
	return arr1
}

// 解法二 暴力遍历排序，时间复杂度高
func getAllElements1(root1 *TreeNode, root2 *TreeNode) []int {
	arr := []int{}
	arr = append(arr, preorderTraversal(root1)...)
	arr = append(arr, preorderTraversal(root2)...)
	sort.Ints(arr)
	return arr
}
