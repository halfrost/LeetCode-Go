package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法一 DFS
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res, maxHeight := 0, -1
	findBottomLeftValueDFS(root, 0, &res, &maxHeight)
	return res
}

func findBottomLeftValueDFS(root *TreeNode, curHeight int, res, maxHeight *int) {
	if curHeight > *maxHeight && root.Left == nil && root.Right == nil {
		*maxHeight = curHeight
		*res = root.Val
	}
	if root.Left != nil {
		findBottomLeftValueDFS(root.Left, curHeight+1, res, maxHeight)
	}
	if root.Right != nil {
		findBottomLeftValueDFS(root.Right, curHeight+1, res, maxHeight)
	}
}

// 解法二 BFS
func findBottomLeftValue1(root *TreeNode) int {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		next := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		if len(next) == 0 {
			return queue[0].Val
		}
		queue = next
	}
	return 0
}
