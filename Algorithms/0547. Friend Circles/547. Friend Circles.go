package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

// 解法一 并查集

func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}
	uf := template.UnionFind{}
	uf.Init(n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.TotalCount()
}

// 解法二 FloodFill DFS 暴力解法
func findCircleNum1(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	visited := make([]bool, len(M))
	res := 0
	for i := range M {
		if !visited[i] {
			dfs547(M, i, visited)
			res++
		}
	}
	return res
}

func dfs547(M [][]int, cur int, visited []bool) {
	visited[cur] = true
	for j := 0; j < len(M[cur]); j++ {
		if !visited[j] && M[cur][j] == 1 {
			dfs547(M, j, visited)
		}
	}
}
