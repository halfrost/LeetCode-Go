package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	queue, res := []*TreeNode{}, 0
	queue = append(queue, &TreeNode{0, root.Left, root.Right})

	for len(queue) != 0 {
		var left, right *int
		// 这里需要注意，先保存 queue 的个数，相当于拿到此层的总个数
		qLen := len(queue)
		// 这里循环不要写 i < len(queue)，因为每次循环 queue 的长度都在变小
		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				// 根据满二叉树父子节点的关系，得到下一层节点在本层的编号
				newVal := node.Val * 2
				queue = append(queue, &TreeNode{newVal, node.Left.Left, node.Left.Right})
				if left == nil || *left > newVal {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
			}
			if node.Right != nil {
				// 根据满二叉树父子节点的关系，得到下一层节点在本层的编号
				newVal := node.Val*2 + 1
				queue = append(queue, &TreeNode{newVal, node.Right.Left, node.Right.Right})
				if left == nil || *left > newVal {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
			}
		}
		switch {
		// 某层只有一个点，那么此层宽度为 1
		case left != nil && right == nil, left == nil && right != nil:
			res = max(res, 1)
		// 某层只有两个点，那么此层宽度为两点之间的距离
		case left != nil && right != nil:
			res = max(res, *right-*left+1)
		}
	}
	return res
}
