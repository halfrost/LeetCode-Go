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
	inorder_(root, k, &count, &res)
	return res
}

func inorder_(node *TreeNode, k int, count *int, ans *int) {
	if node != nil {
		inorder_(node.Left, k, count, ans)
		*count++
		if *count == k {
			*ans = node.Val
			return
		}
		inorder_(node.Right, k, count, ans)
	}
}
