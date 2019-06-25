package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum_III(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := findPath_(root, sum)
	res += pathSum_III(root.Left, sum)
	res += pathSum_III(root.Right, sum)
	return res
}

// 寻找包含 root 这个结点，且和为 sum 的路径
func findPath_(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res += 1
	}
	res += findPath_(root.Left, sum-root.Val)
	res += findPath_(root.Right, sum-root.Val)
	return res
}
