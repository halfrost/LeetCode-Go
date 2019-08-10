package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSameTree(s, t) {
		return true
	}
	if s == nil {
		return false
	}
	if isSubtree(s.Left, t) || isSubtree(s.Right, t) {
		return true
	}
	return false
}
