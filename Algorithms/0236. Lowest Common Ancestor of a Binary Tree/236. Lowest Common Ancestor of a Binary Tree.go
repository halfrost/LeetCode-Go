package leetcode

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor_(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor_(root.Left, p, q)
	right := lowestCommonAncestor_(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
