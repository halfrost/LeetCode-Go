package leetcode

func shiftGrid(grid [][]int, k int) [][]int {
	x, y := len(grid[0]), len(grid)
	newGrid := make([][]int, y)
	for i := 0; i < y; i++ {
		newGrid[i] = make([]int, x)
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			ny := (k / x) + i
			if (j + (k % x)) >= x {
				ny++
			}
			newGrid[ny%y][(j+(k%x))%x] = grid[i][j]
		}
	}
	return newGrid
}
