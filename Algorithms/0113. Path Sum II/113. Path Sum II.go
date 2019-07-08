package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法一
func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	slice = findPath(root, sum, slice, []int(nil))
	return slice
}

func findPath(n *TreeNode, sum int, slice [][]int, stack []int) [][]int {
	if n == nil {
		return slice
	}
	sum -= n.Val
	stack = append(stack, n.Val)
	if sum == 0 && n.Left == nil && n.Right == nil {
		slice = append(slice, append([]int{}, stack...))
		stack = stack[:len(stack)-1]
	}
	slice = findPath(n.Left, sum, slice, stack)
	slice = findPath(n.Right, sum, slice, stack)
	return slice
}

// 解法二
func pathSum1(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return [][]int{[]int{root.Val}}
		}
	}
	path, res := []int{}, [][]int{}
	tmpLeft := pathSum(root.Left, sum-root.Val)
	path = append(path, root.Val)
	if len(tmpLeft) > 0 {
		for i := 0; i < len(tmpLeft); i++ {
			tmpLeft[i] = append(path, tmpLeft[i]...)
		}
		res = append(res, tmpLeft...)
	}
	path = []int{}
	tmpRight := pathSum(root.Right, sum-root.Val)
	path = append(path, root.Val)

	if len(tmpRight) > 0 {
		for i := 0; i < len(tmpRight); i++ {
			tmpRight[i] = append(path, tmpRight[i]...)
		}
		res = append(res, tmpRight...)
	}
	return res
}
