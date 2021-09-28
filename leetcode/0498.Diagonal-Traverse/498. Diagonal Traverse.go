package leetcode

// 解法一
func findDiagonalOrder1(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	row, col, dir, i, x, y, d := len(matrix), len(matrix[0]), [2][2]int{
		{-1, 1},
		{1, -1},
	}, 0, 0, 0, 0
	total := row * col
	res := make([]int, total)
	for i < total {
		for x >= 0 && x < row && y >= 0 && y < col {
			res[i] = matrix[x][y]
			i++
			x += dir[d][0]
			y += dir[d][1]
		}
		d = (d + 1) % 2
		if x == row {
			x--
			y += 2
		}
		if y == col {
			y--
			x += 2
		}
		if x < 0 {
			x = 0
		}
		if y < 0 {
			y = 0
		}
	}
	return res
}

// 解法二
func findDiagonalOrder2(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}
	// dir = 0 代表从右上到左下的方向， dir = 1 代表从左下到右上的方向 dir = -1 代表上一次转变了方向
	m, n, i, j, dir, res := len(matrix), len(matrix[0]), 0, 0, 0, []int{}
	for index := 0; index < m*n; index++ {
		if dir == -1 {
			if (i == 0 && j < n-1) || (j == n-1) { // 上边界和右边界
				i++
				if j > 0 {
					j--
				}
				dir = 0
				addTraverse(matrix, i, j, &res)
				continue
			}
			if (j == 0 && i < m-1) || (i == m-1) { // 左边界和下边界
				if i > 0 {
					i--
				}
				j++
				dir = 1
				addTraverse(matrix, i, j, &res)
				continue
			}
		}
		if i == 0 && j == 0 {
			res = append(res, matrix[i][j])
			if j < n-1 {
				j++
				dir = -1
				addTraverse(matrix, i, j, &res)
				continue
			} else {
				if i < m-1 {
					i++
					dir = -1
					addTraverse(matrix, i, j, &res)
					continue
				}
			}
		}
		if i == 0 && j < n-1 { // 上边界
			if j < n-1 {
				j++
				dir = -1
				addTraverse(matrix, i, j, &res)
				continue
			}
		}
		if j == 0 && i < m-1 { // 左边界
			if i < m-1 {
				i++
				dir = -1
				addTraverse(matrix, i, j, &res)
				continue
			}
		}
		if j == n-1 { // 右边界
			if i < m-1 {
				i++
				dir = -1
				addTraverse(matrix, i, j, &res)
				continue
			}
		}
		if i == m-1 { // 下边界
			j++
			dir = -1
			addTraverse(matrix, i, j, &res)
			continue
		}
		if dir == 1 {
			i--
			j++
			addTraverse(matrix, i, j, &res)
			continue
		}
		if dir == 0 {
			i++
			j--
			addTraverse(matrix, i, j, &res)
			continue
		}
	}
	return res
}

func addTraverse(matrix [][]int, i, j int, res *[]int) {
	if i >= 0 && i <= len(matrix)-1 && j >= 0 && j <= len(matrix[0])-1 {
		*res = append(*res, matrix[i][j])
	}
}

// 解法三
func findDiagonalOrder(mat [][]int) []int {
	forUp := true // true: 斜上角  false: 斜下角
	// 初始化位置
	idxX := 0
	idxY := 0
	var ret []int
	x := len(mat[0])
	y := len(mat)
	for idxX < x && idxY < y {
		ret = append(ret, mat[idxY][idxX])
		if forUp {
			if idxX+1 < x && idxY-1 >= 0 { // 斜上角 有值
				idxX += 1
				idxY -= 1
			} else if idxX+1 < x { // 斜上角 无值 且 x轴右侧有值 ，优先水平向右移动
				idxX += 1
				forUp = false
			} else { // 否则 ， 垂直向下移动
				idxY += 1
				forUp = false
			}
		} else {
			if idxX-1 >= 0 && idxY+1 < y { // 斜下角 有值
				idxX -= 1
				idxY += 1
			} else if idxY+1 < y { // 斜下角 无值  且 Y轴下侧有值，优先垂直向下移动
				idxY += 1
				forUp = true
			} else { // 否则 ， 水平向右移动
				idxX += 1
				forUp = true
			}
		}
	}
	return ret
}
