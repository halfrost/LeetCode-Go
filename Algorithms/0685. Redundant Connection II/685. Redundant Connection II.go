package leetcode

func findRedundantDirectedConnection(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	parent, candidate1, candidate2 := make([]int, len(edges)+1), []int{}, []int{}
	for _, edge := range edges {
		if parent[edge[1]] == 0 {
			parent[edge[1]] = edge[0]
		} else { // 如果一个节点已经有父亲节点了，说明入度已经有 1 了，再来一条边，入度为 2 ，那么跳过新来的这条边 candidate2，并记录下和这条边冲突的边 candidate1
			candidate1 = append(candidate1, parent[edge[1]])
			candidate1 = append(candidate1, edge[1])
			candidate2 = append(candidate2, edge[0])
			candidate2 = append(candidate2, edge[1])
			edge[1] = 0 // 做标记，后面再扫到这条边以后可以直接跳过
		}
	}
	for i := 1; i <= len(edges); i++ {
		parent[i] = i
	}
	for _, edge := range edges {
		if edge[1] == 0 { // 跳过 candidate2 这条边
			continue
		}
		u, v := edge[0], edge[1]
		pu := findRoot(&parent, u)
		if pu == v { // 发现有环
			if len(candidate1) == 0 { // 如果没有出现入度为 2 的情况，那么对应情况 1，就删除这条边
				return edge
			}
			return candidate1 // 出现环并且有入度为 2 的情况，说明 candidate1 是答案
		}
		parent[v] = pu // 没有发现环，继续合并
	}
	return candidate2 // 当最后什么都没有发生，则 candidate2 是答案
}

func findRoot(parent *[]int, k int) int {
	if (*parent)[k] != k {
		(*parent)[k] = findRoot(parent, (*parent)[k])
	}
	return (*parent)[k]
}
