package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type status int

const (
	isLeaf status = iota
	parentofLeaf
	isMonitoredWithoutCamera
)

func minCameraCover(root *TreeNode) int {
	res := 0
	if minCameraCoverDFS(root, &res) == isLeaf {
		res++
	}
	return res
}

func minCameraCoverDFS(root *TreeNode, res *int) status {
	if root == nil {
		return 2
	}
	left, right := minCameraCoverDFS(root.Left, res), minCameraCoverDFS(root.Right, res)
	if left == isLeaf || right == isLeaf {
		*res++
		return parentofLeaf
	} else if left == parentofLeaf || right == parentofLeaf {
		return isMonitoredWithoutCamera
	} else {
		return isLeaf
	}
}
