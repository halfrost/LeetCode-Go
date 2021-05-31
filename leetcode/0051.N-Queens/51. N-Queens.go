package leetcode

// 解法一 DFS
func solveNQueens(n int) [][]string {
	col, dia1, dia2, row, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, [][]string{}
	putQueen(n, 0, &col, &dia1, &dia2, &row, &res)
	return res
}

// 尝试在一个n皇后问题中, 摆放第index行的皇后位置
func putQueen(n, index int, col, dia1, dia2 *[]bool, row *[]int, res *[][]string) {
	if index == n {
		*res = append(*res, generateBoard(n, row))
		return
	}
	for i := 0; i < n; i++ {
		// 尝试将第index行的皇后摆放在第i列
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i] = true
			(*dia2)[index-i+n-1] = true
			putQueen(n, index+1, col, dia1, dia2, row, res)
			(*col)[i] = false
			(*dia1)[index+i] = false
			(*dia2)[index-i+n-1] = false
			*row = (*row)[:len(*row)-1]
		}
	}
	return
}

func generateBoard(n int, row *[]int) []string {
	board := []string{}
	res := ""
	for i := 0; i < n; i++ {
		res += "."
	}
	for i := 0; i < n; i++ {
		board = append(board, res)
	}
	for i := 0; i < n; i++ {
		tmp := []byte(board[i])
		tmp[(*row)[i]] = 'Q'
		board[i] = string(tmp)
	}
	return board
}

// 解法二 二进制操作法 Signed-off-by: Hanlin Shi shihanlin9@gmail.com
func solveNQueens2(n int) (res [][]string) {
	placements := make([]string, n)
	for i := range placements {
		buf := make([]byte, n)
		for j := range placements {
			if i == j {
				buf[j] = 'Q'
			} else {
				buf[j] = '.'
			}
		}
		placements[i] = string(buf)
	}
	var construct func(prev []int)
	construct = func(prev []int) {
		if len(prev) == n {
			plan := make([]string, n)
			for i := 0; i < n; i++ {
				plan[i] = placements[prev[i]]
			}
			res = append(res, plan)
			return
		}
		occupied := 0
		for i := range prev {
			dist := len(prev) - i
			bit := 1 << prev[i]
			occupied |= bit | bit<<dist | bit>>dist
		}
		prev = append(prev, -1)
		for i := 0; i < n; i++ {
			if (occupied>>i)&1 != 0 {
				continue
			}
			prev[len(prev)-1] = i
			construct(prev)
		}
	}
	construct(make([]int, 0, n))
	return
}
