package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum, res, count, sum := 1, 0, []float64{}, 1, 0
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
			sum += node.Val
			queue = queue[1:]
		}
		if curNum == 0 {
			res = append(res, float64(sum)/float64(count))
			curNum, count, nextLevelNum, sum = nextLevelNum, nextLevelNum, 0, 0
		}
	}
	return res
}
