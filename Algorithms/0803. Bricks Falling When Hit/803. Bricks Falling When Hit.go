package leetcode

func hitBricks(grid [][]int, hits [][]int) []int {
	if len(hits) == 0 {
		return []int{}
	}
	uf, m, n, res, oriCount := UnionFindCount{}, len(grid), len(grid[0]), make([]int, len(hits)), 0
	uf.init(m*n + 1)
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
	oriCount = uf.count[uf.find(m*n)]
	for i := len(hits) - 1; i >= 0; i-- {
		if grid[hits[i][0]][hits[i][1]] == 2 {
			grid[hits[i][0]][hits[i][1]] = 1
			getUnionFindFromGrid(grid, hits[i][0], hits[i][1], uf)
		}
		nowCount := uf.count[uf.find(m*n)]
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

func getUnionFindFromGrid(grid [][]int, x, y int, uf UnionFindCount) {
	m, n := len(grid), len(grid[0])
	if x == 0 {
		uf.union(m*n, x*n+y)
	}
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInGrid(grid, nx, ny) && grid[nx][ny] == 1 {
			uf.union(nx*n+ny, x*n+y)
		}
	}
}

// UnionFindCount define
type UnionFindCount struct {
	parent, count []int
}

func (uf *UnionFindCount) init(n int) {
	uf.parent = make([]int, n)
	uf.count = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
		uf.count[i] = 1
	}
}

func (uf *UnionFindCount) find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	return root
}

// 不进行秩压缩，时间复杂度爆炸，太高了
// func (uf *UnionFindCount) union(p, q int) {
// 	proot := uf.find(p)
// 	qroot := uf.find(q)
// 	if proot == qroot {
// 		return
// 	}
// 	if proot != qroot {
// 		uf.parent[proot] = qroot
// 		uf.count[qroot] += uf.count[proot]
// 	}
// }

func (uf *UnionFindCount) union(p, q int) {
	proot := uf.find(p)
	qroot := uf.find(q)
	if proot == qroot {
		return
	}
	if proot == len(uf.parent)-1 {
		//proot is root
	} else if qroot == len(uf.parent)-1 {
		// qroot is root, always attach to root
		proot, qroot = qroot, proot
	} else if uf.count[qroot] > uf.count[proot] {
		proot, qroot = qroot, proot
	}

	//set relation[0] as parent
	uf.parent[qroot] = proot
	uf.count[proot] += uf.count[qroot]
}
