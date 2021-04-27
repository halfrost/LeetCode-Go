package leetcode

//  Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// 解法一 非递归
func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, r.Val)
		tmp := []*Node{}
		for _, v := range r.Children {
			tmp = append([]*Node{v}, tmp...) // 逆序存点
		}
		stack = append(stack, tmp...)
	}
	return res
}

// 解法二 递归
func preorder1(root *Node) []int {
	res := []int{}
	preorderdfs(root, &res)
	return res
}

func preorderdfs(root *Node, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		for i := 0; i < len(root.Children); i++ {
			preorderdfs(root.Children[i], res)
		}
	}
}
