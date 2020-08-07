package leetcode

func eventualSafeNodes(graph [][]int) []int {
	res, color := []int{}, make([]int, len(graph))
	for i := range graph {
		if dfsEventualSafeNodes(graph, i, color) {
			res = append(res, i)
		}
	}
	return res
}

// colors: WHITE 0, GRAY 1, BLACK 2;
func dfsEventualSafeNodes(graph [][]int, idx int, color []int) bool {
	if color[idx] > 0 {
		return color[idx] == 2
	}
	color[idx] = 1
	for i := range graph[idx] {
		if !dfsEventualSafeNodes(graph, graph[idx][i], color) {
			return false
		}
	}
	color[idx] = 2
	return true
}
