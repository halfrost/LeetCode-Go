package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法一 递归
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	levelX, levelY := findLevel(root, x, 1), findLevel(root, y, 1)
	if levelX != levelY {
		return false
	}
	return !haveSameParents(root, x, y)
}

func findLevel(root *TreeNode, x, level int) int {
	if root == nil {
		return 0
	}
	if root.Val != x {
		leftLevel, rightLevel := findLevel(root.Left, x, level+1), findLevel(root.Right, x, level+1)
		if leftLevel == 0 {
			return rightLevel
		}
		return leftLevel
	}
	return level
}

func haveSameParents(root *TreeNode, x, y int) bool {
	if root == nil {
		return false
	}
	if (root.Left != nil && root.Right != nil && root.Left.Val == x && root.Right.Val == y) ||
		(root.Left != nil && root.Right != nil && root.Left.Val == y && root.Right.Val == x) {
		return true
	}
	return haveSameParents(root.Left, x, y) || haveSameParents(root.Right, x, y)
}

// 解法二 BFS
type mark struct {
	prev  int
	depth int
}

func isCousinsBFS(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}
	visited := [101]*mark{}
	visited[root.Val] = &mark{prev: -1, depth: 1}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		depth := visited[node.Val].depth
		if node.Left != nil {
			visited[node.Left.Val] = &mark{prev: node.Val, depth: depth + 1}
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			visited[node.Right.Val] = &mark{prev: node.Val, depth: depth + 1}
			queue = append(queue, node.Right)
		}
	}
	if visited[x] == nil || visited[y] == nil {
		return false
	}
	if visited[x].depth == visited[y].depth && visited[x].prev != visited[y].prev {
		return true
	}
	return false
}

// 解法三 DFS
func isCousinsDFS(root *TreeNode, x int, y int) bool {
	var depth1, depth2, parent1, parent2 int
	dfsCousins(root, x, 0, -1, &parent1, &depth1)
	dfsCousins(root, y, 0, -1, &parent2, &depth2)
	return depth1 > 1 && depth1 == depth2 && parent1 != parent2
}

func dfsCousins(root *TreeNode, val, depth, last int, parent, res *int) {
	if root == nil {
		return
	}
	if root.Val == val {
		*res = depth
		*parent = last
		return
	}
	depth++
	dfsCousins(root.Left, val, depth, root.Val, parent, res)
	dfsCousins(root.Right, val, depth, root.Val, parent, res)
}
