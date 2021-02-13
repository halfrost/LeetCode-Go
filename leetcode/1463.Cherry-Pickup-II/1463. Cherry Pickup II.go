package leetcode

func cherryPickup(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([][]int, cols)
		for j := 0; j < cols; j++ {
			dp[i][j] = make([]int, cols)
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j <= i && j < cols; j++ {
			for k := cols - 1; k >= cols-1-i && k >= 0; k-- {
				max := 0
				for a := j - 1; a <= j+1; a++ {
					for b := k - 1; b <= k+1; b++ {
						sum := isInBoard(dp, i-1, a, b)
						if a == b && i > 0 && a >= 0 && a < cols {
							sum -= grid[i-1][a]
						}
						if sum > max {
							max = sum
						}
					}
				}
				if j == k {
					max += grid[i][j]
				} else {
					max += grid[i][j] + grid[i][k]
				}
				dp[i][j][k] = max
			}
		}
	}
	count := 0
	for j := 0; j < cols && j < rows; j++ {
		for k := cols - 1; k >= 0 && k >= cols-rows; k-- {
			if dp[rows-1][j][k] > count {
				count = dp[rows-1][j][k]
			}
		}
	}
	return count
}

func isInBoard(dp [][][]int, i, j, k int) int {
	if i < 0 || j < 0 || j >= len(dp[0]) || k < 0 || k >= len(dp[0]) {
		return 0
	}
	return dp[i][j][k]
}
