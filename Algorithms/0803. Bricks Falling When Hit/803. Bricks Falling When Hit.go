package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

func hitBricks(grid [][]int, hits [][]int) []int {
	if len(hits) == 0 {
		return []int{}
	}
	uf, m, n, res, oriCount := template.UnionFindCount{}, len(grid), len(grid[0]), make([]int, len(hits)), 0
	uf.Init(m*n + 1)
	// 先将要打掉的砖块染色
	for _, hit := range hits {
		if grid[hit[0]][hit[1]] == 1 {
			grid[hit[0]][hit[1]] = 2
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				getUnionFindFromGrid(grid, i, j, uf)
			}
		}
	}
	oriCount = uf.Count()[uf.Find(m*n)]
	for i := len(hits) - 1; i >= 0; i-- {
		if grid[hits[i][0]][hits[i][1]] == 2 {
			grid[hits[i][0]][hits[i][1]] = 1
			getUnionFindFromGrid(grid, hits[i][0], hits[i][1], uf)
		}
		nowCount := uf.Count()[uf.Find(m*n)]
		if nowCount-oriCount > 0 {
			res[i] = nowCount - oriCount - 1
		} else {
			res[i] = 0
		}
		oriCount = nowCount
	}
	return res
}

func isInGrid(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func getUnionFindFromGrid(grid [][]int, x, y int, uf template.UnionFindCount) {
	m, n := len(grid), len(grid[0])
	if x == 0 {
		uf.Union(m*n, x*n+y)
	}
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInGrid(grid, nx, ny) && grid[nx][ny] == 1 {
			uf.Union(nx*n+ny, x*n+y)
		}
	}
}
