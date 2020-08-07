package leetcode

// 解法一 暴力法
func oddCells(n int, m int, indices [][]int) int {
	matrix, res := make([][]int, n), 0
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	for _, indice := range indices {
		for i := 0; i < m; i++ {
			matrix[indice[0]][i]++
		}
		for j := 0; j < n; j++ {
			matrix[j][indice[1]]++
		}
	}
	for _, m := range matrix {
		for _, v := range m {
			if v&1 == 1 {
				res++
			}
		}
	}
	return res
}

// 解法二 暴力法
func oddCells1(n int, m int, indices [][]int) int {
	rows, cols, count := make([]int, n), make([]int, m), 0
	for _, pair := range indices {
		rows[pair[0]]++
		cols[pair[1]]++
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (rows[i]+cols[j])%2 == 1 {
				count++
			}
		}
	}
	return count
}
