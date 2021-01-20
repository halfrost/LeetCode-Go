package leetcode

// DFS 染色，1 是红色，0 是绿色，-1 是未染色
func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	for i := range colors {
		colors[i] = -1
	}
	for i := range graph {
		if !dfs(i, graph, colors, -1) {
			return false
		}
	}
	return true
}

func dfs(n int, graph [][]int, colors []int, parentCol int) bool {
	if colors[n] == -1 {
		if parentCol == 1 {
			colors[n] = 0
		} else {
			colors[n] = 1
		}
	} else if colors[n] == parentCol {
		return false
	} else if colors[n] != parentCol {
		return true
	}
	for _, c := range graph[n] {
		if !dfs(c, graph, colors, colors[n]) {
			return false
		}
	}
	return true
}
