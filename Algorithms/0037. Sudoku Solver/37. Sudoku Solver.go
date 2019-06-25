package leetcode

type position struct {
	x int
	y int
}

func solveSudoku(board [][]byte) {
	pos, find := []position{}, false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				pos = append(pos, position{x: i, y: j})
			}
		}
	}
	putSudoku(&board, pos, 0, &find)
}

func putSudoku(board *[][]byte, pos []position, index int, succ *bool) {
	if *succ == true {
		return
	}
	if index == len(pos) {
		*succ = true
		return
	}
	for i := 1; i < 10; i++ {
		if checkSudoku(board, pos[index], i) && !*succ {
			(*board)[pos[index].x][pos[index].y] = byte(i) + '0'
			putSudoku(board, pos, index+1, succ)
			if *succ == true {
				return
			}
			(*board)[pos[index].x][pos[index].y] = '.'
		}
	}
}

func checkSudoku(board *[][]byte, pos position, val int) bool {
	// 判断横行是否有重复数字
	for i := 0; i < len((*board)[0]); i++ {
		if (*board)[pos.x][i] != '.' && int((*board)[pos.x][i]-'0') == val {
			return false
		}
	}
	// 判断竖行是否有重复数字
	for i := 0; i < len((*board)); i++ {
		if (*board)[i][pos.y] != '.' && int((*board)[i][pos.y]-'0') == val {
			return false
		}
	}
	// 判断九宫格是否有重复数字
	posx, posy := pos.x-pos.x%3, pos.y-pos.y%3
	for i := posx; i < posx+3; i++ {
		for j := posy; j < posy+3; j++ {
			if (*board)[i][j] != '.' && int((*board)[i][j]-'0') == val {
				return false
			}
		}
	}
	return true
}
