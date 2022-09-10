package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

// 解法一 DFS + 二分
func swimInWater(grid [][]int) int {
	row, col, flags, minWait, maxWait := len(grid), len(grid[0]), make([][]int, len(grid)), 0, 0
	for i, row := range grid {
		flags[i] = make([]int, len(row))
		for j := 0; j < col; j++ {
			flags[i][j] = -1
			if row[j] > maxWait {
				maxWait = row[j]
			}
		}
	}
	for minWait < maxWait {
		midWait := (minWait + maxWait) / 2
		addFlags(grid, flags, midWait, 0, 0)
		if flags[row-1][col-1] == midWait {
			maxWait = midWait
		} else {
			minWait = midWait + 1
		}
	}
	return minWait
}

func addFlags(grid [][]int, flags [][]int, flag int, row int, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return
	}
	if grid[row][col] > flag || flags[row][col] == flag {
		return
	}
	flags[row][col] = flag
	addFlags(grid, flags, flag, row-1, col)
	addFlags(grid, flags, flag, row+1, col)
	addFlags(grid, flags, flag, row, col-1)
	addFlags(grid, flags, flag, row, col+1)
}

// 解法二 并查集(并不是此题的最优解)
func swimInWater1(grid [][]int) int {
	n, uf, res := len(grid), template.UnionFind{}, 0
	uf.Init(n * n)
	for uf.Find(0) != uf.Find(n*n-1) {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] > res {
					continue
				}
				if i < n-1 && grid[i+1][j] <= res {
					uf.Union(i*n+j, i*n+j+n)
				}
				if j < n-1 && grid[i][j+1] <= res {
					uf.Union(i*n+j, i*n+j+1)
				}
			}
		}
		res++
	}
	return res - 1
}
