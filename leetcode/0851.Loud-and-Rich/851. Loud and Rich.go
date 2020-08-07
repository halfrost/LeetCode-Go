package leetcode

func loudAndRich(richer [][]int, quiet []int) []int {
	edges := make([][]int, len(quiet))
	for i := range edges {
		edges[i] = []int{}
	}
	indegrees := make([]int, len(quiet))
	for _, edge := range richer {
		n1, n2 := edge[0], edge[1]
		edges[n1] = append(edges[n1], n2)
		indegrees[n2]++
	}
	res := make([]int, len(quiet))
	for i := range res {
		res[i] = i
	}
	queue := []int{}
	for i, v := range indegrees {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		nexts := []int{}
		for _, n1 := range queue {
			for _, n2 := range edges[n1] {
				indegrees[n2]--
				if quiet[res[n2]] > quiet[res[n1]] {
					res[n2] = res[n1]
				}
				if indegrees[n2] == 0 {
					nexts = append(nexts, n2)
				}
			}
		}
		queue = nexts
	}
	return res
}
