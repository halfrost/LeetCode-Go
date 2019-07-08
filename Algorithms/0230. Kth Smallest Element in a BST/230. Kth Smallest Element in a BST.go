package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0
	inorder230(root, k, &count, &res)
	return res
}

func inorder230(node *TreeNode, k int, count *int, ans *int) {
	if node != nil {
		inorder230(node.Left, k, count, ans)
		*count++
		if *count == k {
			*ans = node.Val
			return
		}
		inorder230(node.Right, k, count, ans)
	}
}
