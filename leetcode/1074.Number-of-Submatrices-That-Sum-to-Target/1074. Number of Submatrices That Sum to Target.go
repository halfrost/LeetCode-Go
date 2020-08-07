package leetcode

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n, res := len(matrix), len(matrix[0]), 0
	for row := range matrix {
		for col := 1; col < len(matrix[row]); col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			counterMap, sum := make(map[int]int, m), 0
			counterMap[0] = 1 // 题目保证一定有解，所以这里初始化是 1
			for row := 0; row < m; row++ {
				if i > 0 {
					sum += matrix[row][j] - matrix[row][i-1]
				} else {
					sum += matrix[row][j]
				}
				res += counterMap[sum-target]
				counterMap[sum]++
			}
		}
	}
	return res
}

// 暴力解法 O(n^4)
func numSubmatrixSumTarget1(matrix [][]int, target int) int {
	m, n, res, sum := len(matrix), len(matrix[0]), 0, 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			counterMap := map[int]int{}
			counterMap[0] = 1 // 题目保证一定有解，所以这里初始化是 1
			sum = 0
			for row := 0; row < m; row++ {
				for k := i; k <= j; k++ {
					sum += matrix[row][k]
				}
				res += counterMap[sum-target]
				counterMap[sum]++
			}
		}
	}
	return res
}

// 暴力解法超时！ O(n^6)
func numSubmatrixSumTarget2(matrix [][]int, target int) int {
	res := 0
	for startx := 0; startx < len(matrix); startx++ {
		for starty := 0; starty < len(matrix[startx]); starty++ {
			for endx := startx; endx < len(matrix); endx++ {
				for endy := starty; endy < len(matrix[startx]); endy++ {
					if sumSubmatrix(matrix, startx, starty, endx, endy) == target {
						//fmt.Printf("startx = %v, starty = %v, endx = %v, endy = %v\n", startx, starty, endx, endy)
						res++
					}
				}
			}
		}
	}
	return res
}

func sumSubmatrix(matrix [][]int, startx, starty, endx, endy int) int {
	sum := 0
	for i := startx; i <= endx; i++ {
		for j := starty; j <= endy; j++ {
			sum += matrix[i][j]
		}
	}
	return sum
}
