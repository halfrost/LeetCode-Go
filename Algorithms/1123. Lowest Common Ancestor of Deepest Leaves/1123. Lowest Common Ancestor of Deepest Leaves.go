package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	lca, maxLevel := &TreeNode{}, 0
	lcaDeepestLeavesDFS(&lca, &maxLevel, 0, root)
	return lca
}

func lcaDeepestLeavesDFS(lca **TreeNode, maxLevel *int, depth int, root *TreeNode) int {
	*maxLevel = max(*maxLevel, depth)
	if root == nil {
		return depth
	}
	depthLeft := lcaDeepestLeavesDFS(lca, maxLevel, depth+1, root.Left)
	depthRight := lcaDeepestLeavesDFS(lca, maxLevel, depth+1, root.Right)
	if depthLeft == *maxLevel && depthRight == *maxLevel {
		*lca = root
	}
	return max(depthLeft, depthRight)
}
