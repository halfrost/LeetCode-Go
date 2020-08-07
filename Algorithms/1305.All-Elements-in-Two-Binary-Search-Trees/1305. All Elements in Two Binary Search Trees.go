package leetcode

import (
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

// 解法一 合并排序
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := inorderTraversal(root1)
	arr2 := inorderTraversal(root2)
	arr1 = append(arr1, make([]int, len(arr2))...)
	merge(arr1, len(arr1)-len(arr2), arr2, len(arr2))
	return arr1
}

// this is 94 solution
func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}

// this is 88 solution
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	// 这里不需要，因为测试数据考虑到了第一个数组的空间问题
	// for index := 0; index < n; index++ {
	// 	nums1 = append(nums1, nums2[index])
	// }
	i := m - 1
	j := n - 1
	k := m + n - 1
	// 从后面往前放，只需要循环一次即可
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
}

// 解法二 暴力遍历排序，时间复杂度高
func getAllElements1(root1 *TreeNode, root2 *TreeNode) []int {
	arr := []int{}
	arr = append(arr, preorderTraversal(root1)...)
	arr = append(arr, preorderTraversal(root2)...)
	sort.Ints(arr)
	return arr
}

// this is 144 solution
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		tmp := preorderTraversal(root.Left)
		for _, t := range tmp {
			res = append(res, t)
		}
		tmp = preorderTraversal(root.Right)
		for _, t := range tmp {
			res = append(res, t)
		}
	}
	return res
}
