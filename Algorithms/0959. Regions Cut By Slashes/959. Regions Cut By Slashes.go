package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

func regionsBySlashes(grid []string) int {
	size := len(grid)
	uf := template.UnionFind{}
	uf.Init(4 * size * size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			switch grid[i][j] {
			case '\\':
				uf.Union(getFaceIdx(size, i, j, 0), getFaceIdx(size, i, j, 1))
				uf.Union(getFaceIdx(size, i, j, 2), getFaceIdx(size, i, j, 3))
			case '/':
				uf.Union(getFaceIdx(size, i, j, 0), getFaceIdx(size, i, j, 3))
				uf.Union(getFaceIdx(size, i, j, 2), getFaceIdx(size, i, j, 1))
			case ' ':
				uf.Union(getFaceIdx(size, i, j, 0), getFaceIdx(size, i, j, 1))
				uf.Union(getFaceIdx(size, i, j, 2), getFaceIdx(size, i, j, 1))
				uf.Union(getFaceIdx(size, i, j, 2), getFaceIdx(size, i, j, 3))
			}
			if i < size-1 {
				uf.Union(getFaceIdx(size, i, j, 2), getFaceIdx(size, i+1, j, 0))
			}
			if j < size-1 {
				uf.Union(getFaceIdx(size, i, j, 1), getFaceIdx(size, i, j+1, 3))
			}
		}
	}
	count := 0
	for i := 0; i < 4*size*size; i++ {
		if uf.Find(i) == i {
			count++
		}
	}
	return count
}

func getFaceIdx(size, i, j, k int) int {
	return 4*(i*size+j) + k
}
