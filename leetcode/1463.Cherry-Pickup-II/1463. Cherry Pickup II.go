package leetcode

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	old, new := make([]int, n*n), make([]int, n*n)
	for i := range old {
		old[i] = -0xffffff
	}
	old[n-1] = grid[0][0]+grid[0][n-1]

	// dp
	for k:=1; k<m; k++ {
		for s:=0; s<n*n; s++ {
			new[s] = -0xffffff
			c1, c2 := s/n, s%n
			toadd := grid[k][c1]
			if c1 != c2 {
				toadd += grid[k][c2]
			}
			for _, d1 := range []int{1,0,-1} {
				for _, d2 := range []int{1,0,-1} {
					nc1, nc2 := c1+d1, c2+d2
					if nc1>=0 && nc1<n && nc2>=0 && nc2<n && old[nc1*n+nc2]>=0 {
						new[s] = max(new[s], old[nc1*n+nc2]+toadd)
					}
				}
			}
		}
		old, new = new, old
	}
	allmax := 0
	for _, v := range old {
		if v>allmax {
			allmax = v
		}
	}
	return allmax
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}
