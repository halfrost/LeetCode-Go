package leetcode

func findRedundantDirectedConnection(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	parent, candidate1, candidate2 := make([]int, len(edges)+1), []int{}, []int{}
	for _, edge := range edges {
		if parent[edge[1]] == 0 {
			parent[edge[1]] = edge[0]
		} else {
			candidate1 = append(candidate1, parent[edge[1]])
			candidate1 = append(candidate1, edge[1])
			candidate2 = append(candidate2, edge[0])
			candidate2 = append(candidate2, edge[1])
			edge[1] = 0
		}
	}
	for i := 1; i <= len(edges); i++ {
		parent[i] = i
	}
	for _, edge := range edges {
		if edge[1] == 0 {
			continue
		}
		u, v := edge[0], edge[1]
		pu := findRoot(&parent, u)
		if pu == v {
			if len(candidate1) == 0 {
				return edge
			}
			return candidate1
		}
		parent[v] = pu
	}
	return candidate2
}

func findRoot(parent *[]int, k int) int {
	if (*parent)[k] != k {
		(*parent)[k] = findRoot(parent, (*parent)[k])
	}
	return (*parent)[k]
}
