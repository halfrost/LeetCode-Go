package leetcode

func islandPerimeter(grid [][]int) int {
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if i-1 < 0 || grid[i-1][j] == 0 {
					counter++
				}
				if i+1 >= len(grid) || grid[i+1][j] == 0 {
					counter++
				}
				if j-1 < 0 || grid[i][j-1] == 0 {
					counter++
				}
				if j+1 >= len(grid[0]) || grid[i][j+1] == 0 {
					counter++
				}
			}
		}
	}
	return counter
}
