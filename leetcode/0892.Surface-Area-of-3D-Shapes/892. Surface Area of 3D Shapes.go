package leetcode

func surfaceArea(grid [][]int) int {
	area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			area += grid[i][j]*4 + 2
			// up
			if i > 0 {
				m := min(grid[i][j], grid[i-1][j])
				area -= m
			}
			// down
			if i < len(grid)-1 {
				m := min(grid[i][j], grid[i+1][j])
				area -= m
			}
			// left
			if j > 0 {
				m := min(grid[i][j], grid[i][j-1])
				area -= m
			}
			// right
			if j < len(grid[i])-1 {
				m := min(grid[i][j], grid[i][j+1])
				area -= m
			}
		}
	}
	return area
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
