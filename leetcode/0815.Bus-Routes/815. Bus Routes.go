package leetcode

func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}
	// vertexMap 中 key 是站点，value 是公交车数组，代表这些公交车路线可以到达此站点
	vertexMap, visited, queue, res := map[int][]int{}, make([]bool, len(routes)), []int{}, 0
	for i := 0; i < len(routes); i++ {
		for _, v := range routes[i] {
			tmp := vertexMap[v]
			tmp = append(tmp, i)
			vertexMap[v] = tmp
		}
	}
	queue = append(queue, S)
	for len(queue) > 0 {
		res++
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			vertex := queue[0]
			queue = queue[1:]
			for _, bus := range vertexMap[vertex] {
				if visited[bus] == true {
					continue
				}
				visited[bus] = true
				for _, v := range routes[bus] {
					if v == T {
						return res
					}
					queue = append(queue, v)
				}
			}
		}
	}
	return -1
}
