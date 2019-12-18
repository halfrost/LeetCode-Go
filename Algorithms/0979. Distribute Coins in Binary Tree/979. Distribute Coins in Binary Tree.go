package leetcode

/**
 * Definition for a binary tree root.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	res := 0
	distributeCoinsDFS(root, &res)
	return res
}

func distributeCoinsDFS(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left, right := distributeCoinsDFS(root.Left, res), distributeCoinsDFS(root.Right, res)
	*res += abs(left) + abs(right)
	return left + right + root.Val - 1
}
