package leetcode

func sumOfDistancesInTree(N int, edges [][]int) []int {
	// count[i] 中存储的是以 i 为根节点，所有子树结点和根节点的总数
	tree, visited, count, res := make([][]int, N), make([]bool, N), make([]int, N), make([]int, N)
	for _, e := range edges {
		i, j := e[0], e[1]
		tree[i] = append(tree[i], j)
		tree[j] = append(tree[j], i)
	}
	deepFirstSearch(0, visited, count, res, tree)
	// 重置访问状态，再进行一次 DFS
	visited = make([]bool, N)
	// 进入第二次 DFS 之前，只有 res[0] 里面存的是正确的值，因为第一次 DFS 计算出了以 0 为根节点的所有路径和
	// 第二次 DFS 的目的是把以 0 为根节点的路径和转换成以 n 为根节点的路径和
	deepSecondSearch(0, visited, count, res, tree)

	return res
}

func deepFirstSearch(root int, visited []bool, count, res []int, tree [][]int) {
	visited[root] = true
	for _, n := range tree[root] {
		if visited[n] {
			continue
		}
		deepFirstSearch(n, visited, count, res, tree)
		count[root] += count[n]
		// root 节点到 n 的所有路径和 = 以 n 为根节点到所有子树的路径和 res[n] + root 到 count[n] 中每个节点的个数(root 节点和以 n 为根节点的每个节点都增加一条路径)
		// root 节点和以 n 为根节点的每个节点都增加一条路径 = 以 n 为根节点，子树节点数和根节点数的总和，即 count[n]
		res[root] += res[n] + count[n]
	}
	count[root]++
}

// 从 root 开始，把 root 节点的子节点，依次设置成新的根节点
func deepSecondSearch(root int, visited []bool, count, res []int, tree [][]int) {
	N := len(visited)
	visited[root] = true
	for _, n := range tree[root] {
		if visited[n] {
			continue
		}
		// 根节点从 root 变成 n 后
		// res[root] 存储的是以 root 为根节点到所有节点的路径总长度
		// 1. root 到 n 节点增加的路径长度 = root 节点和以 n 为根节点的每个节点都增加一条路径 = 以 n 为根节点，子树节点数和根节点数的总和，即 count[n]
		// 2. n 到以 n 为根节点的所有子树节点以外的节点增加的路径长度 = n 节点和非 n 为根节点子树的每个节点都增加一条路径 = N - count[n]
		// 所以把根节点从 root 转移到 n，需要增加的路径是上面👆第二步计算的，需要减少的路径是上面👆第一步计算的
		res[n] = res[root] + (N - count[n]) - count[n]
		deepSecondSearch(n, visited, count, res, tree)
	}
}
