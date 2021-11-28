package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + bfs(root)
}

func bfs(root *Node) int {
	var q []*Node
	var depth int
	q = append(q, root.Children...)
	for len(q) != 0 {
		depth++
		length := len(q)
		for length != 0 {
			ele := q[0]
			q = q[1:]
			length--
			if ele != nil && len(ele.Children) != 0 {
				q = append(q, ele.Children...)
			}
		}
	}
	return depth
}
