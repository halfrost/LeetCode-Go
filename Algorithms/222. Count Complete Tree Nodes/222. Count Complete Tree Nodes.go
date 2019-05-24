package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum, res := 1, 0, 1
	for len(queue) != 0 {
		if curNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			queue = queue[1:]
		}
		if curNum == 0 {
			res += nextLevelNum
			curNum = nextLevelNum
			nextLevelNum = 0
		}
	}
	return res
}
