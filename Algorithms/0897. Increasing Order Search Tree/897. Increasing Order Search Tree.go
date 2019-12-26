package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法一 链表思想
func increasingBST(root *TreeNode) *TreeNode {
	var head = &TreeNode{}
	tail := head
	recBST(root, tail)
	return head.Right
}

func recBST(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}
	tail = recBST(root.Left, tail)
	root.Left = nil               // 切断 root 与其 Left 的连接，避免形成环
	tail.Right, tail = root, root // 把 root 接上 tail，并保持 tail 指向尾部
	tail = recBST(root.Right, tail)
	return tail
}

// 解法二 模拟
func increasingBST1(root *TreeNode) *TreeNode {
	list, newRoot := []int{}, &TreeNode{}
	inorder(root, &list)
	if len(list) == 0 {
		return root
	}
	newRoot = &TreeNode{Val: list[0], Left: nil, Right: nil}
	cur := newRoot
	for index := 1; index < len(list); index++ {
		tmp := &TreeNode{Val: list[index], Left: nil, Right: nil}
		cur.Right = tmp
		cur = tmp
	}
	return newRoot
}
