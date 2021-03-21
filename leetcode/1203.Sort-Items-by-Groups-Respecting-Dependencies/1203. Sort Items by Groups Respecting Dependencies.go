package leetcode

// 解法一 拓扑排序版的 DFS
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	groups, inDegrees := make([][]int, n+m), make([]int, n+m)
	for i, g := range group {
		if g > -1 {
			g += n
			groups[g] = append(groups[g], i)
			inDegrees[i]++
		}
	}
	for i, ancestors := range beforeItems {
		gi := group[i]
		if gi == -1 {
			gi = i
		} else {
			gi += n
		}
		for _, ancestor := range ancestors {
			ga := group[ancestor]
			if ga == -1 {
				ga = ancestor
			} else {
				ga += n
			}
			if gi == ga {
				groups[ancestor] = append(groups[ancestor], i)
				inDegrees[i]++
			} else {
				groups[ga] = append(groups[ga], gi)
				inDegrees[gi]++
			}
		}
	}
	res := []int{}
	for i, d := range inDegrees {
		if d == 0 {
			sortItemsDFS(i, n, &res, &inDegrees, &groups)
		}
	}
	if len(res) != n {
		return nil
	}
	return res
}

func sortItemsDFS(i, n int, res, inDegrees *[]int, groups *[][]int) {
	if i < n {
		*res = append(*res, i)
	}
	(*inDegrees)[i] = -1
	for _, ch := range (*groups)[i] {
		if (*inDegrees)[ch]--; (*inDegrees)[ch] == 0 {
			sortItemsDFS(ch, n, res, inDegrees, groups)
		}
	}
}

// 解法二 二维拓扑排序 时间复杂度 O(m+n)，空间复杂度 O(m+n)
func sortItems1(n int, m int, group []int, beforeItems [][]int) []int {
	groupItems, res := map[int][]int{}, []int{}
	for i := 0; i < len(group); i++ {
		if group[i] == -1 {
			group[i] = m + i
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}
	groupGraph, groupDegree, itemGraph, itemDegree := make([][]int, m+n), make([]int, m+n), make([][]int, n), make([]int, n)
	for i := 0; i < len(beforeItems); i++ {
		for j := 0; j < len(beforeItems[i]); j++ {
			if group[beforeItems[i][j]] != group[i] {
				// 不同组项目，确定组间依赖关系
				groupGraph[group[beforeItems[i][j]]] = append(groupGraph[group[beforeItems[i][j]]], group[i])
				groupDegree[group[i]]++
			} else {
				// 同组项目，确定组内依赖关系
				itemGraph[beforeItems[i][j]] = append(itemGraph[beforeItems[i][j]], i)
				itemDegree[i]++
			}
		}
	}
	items := []int{}
	for i := 0; i < m+n; i++ {
		items = append(items, i)
	}
	// 组间拓扑
	groupOrders := topSort(groupGraph, groupDegree, items)
	if len(groupOrders) < len(items) {
		return nil
	}
	for i := 0; i < len(groupOrders); i++ {
		items := groupItems[groupOrders[i]]
		// 组内拓扑
		orders := topSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		res = append(res, orders...)
	}
	return res
}

func topSort(graph [][]int, deg, items []int) (orders []int) {
	q := []int{}
	for _, i := range items {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		orders = append(orders, from)
		for _, to := range graph[from] {
			deg[to]--
			if deg[to] == 0 {
				q = append(q, to)
			}
		}
	}
	return
}
