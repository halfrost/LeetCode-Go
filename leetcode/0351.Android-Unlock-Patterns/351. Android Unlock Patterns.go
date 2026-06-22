package leetcode

// blocker351[a][b] 记录从点 a 滑到点 b 必须先经过的中间点；0 表示两点相邻、无需经过中间点。
// 例如 1->3 必须先经过 2，1->9 必须先经过 5。
var blocker351 = [10][10]int{
	1: {3: 2, 7: 4, 9: 5},
	2: {8: 5},
	3: {1: 2, 7: 5, 9: 6},
	4: {6: 5},
	6: {4: 5},
	7: {1: 4, 3: 5, 9: 8},
	8: {2: 5},
	9: {1: 5, 3: 6, 7: 8},
}

// 解法一 回溯。0 号点是虚拟起点，从它出发可以到任意一个数字。
func numberOfPatterns(m int, n int) int {
	visited := make([]bool, 10)
	return countPatterns351(visited, 0, 0, m, n)
}

// now 是当前所在的点，already 是已经选中的点数，统计长度落在 [m,n] 内的合法图案数。
func countPatterns351(visited []bool, now, already, m, n int) int {
	res := 0
	if already >= m && already <= n {
		res++
	}
	if already >= n { // 已达到最大长度，无需再往下选
		return res
	}
	for next := 1; next <= 9; next++ {
		mid := blocker351[now][next]
		// next 没被用过，且 next 与 now 相邻（mid==0）或中间点已被划过（visited[mid]）
		if !visited[next] && (mid == 0 || visited[mid]) {
			visited[next] = true
			res += countPatterns351(visited, next, already+1, m, n)
			visited[next] = false
		}
	}
	return res
}
