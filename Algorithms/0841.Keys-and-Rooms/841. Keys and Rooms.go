package leetcode

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	visited[0] = true
	dfsVisitAllRooms(rooms, visited, 0)
	return len(rooms) == len(visited)
}

func dfsVisitAllRooms(es [][]int, visited map[int]bool, from int) {
	for _, to := range es[from] {
		if visited[to] {
			continue
		}
		visited[to] = true
		dfsVisitAllRooms(es, visited, to)
	}
}
