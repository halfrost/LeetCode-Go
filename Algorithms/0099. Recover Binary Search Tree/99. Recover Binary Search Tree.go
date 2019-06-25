package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	var prev, target1, target2 *TreeNode
	_, target1, target2 = inOrderTraverse(root, prev, target1, target2)
	if target1 != nil && target2 != nil {
		target1.Val, target2.Val = target2.Val, target1.Val
	}

}

func inOrderTraverse(root, prev, target1, target2 *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if root == nil {
		return prev, target1, target2
	}
	prev, target1, target2 = inOrderTraverse(root.Left, prev, target1, target2)
	if prev != nil && prev.Val > root.Val {
		if target1 == nil {
			target1 = prev
		}
		target2 = root
	}
	prev = root
	prev, target1, target2 = inOrderTraverse(root.Right, prev, target1, target2)
	return prev, target1, target2
}
