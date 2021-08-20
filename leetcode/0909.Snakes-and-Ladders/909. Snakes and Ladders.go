package leetcode

type pair struct {
	id, step int
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	visited := make([]bool, n*n+1)
	queue := []pair{{1, 0}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for i := 1; i <= 6; i++ {
			nxt := p.id + i
			if nxt > n*n { // 超出边界
				break
			}
			r, c := getRowCol(nxt, n) // 得到下一步的行列
			if board[r][c] > 0 {      // 存在蛇或梯子
				nxt = board[r][c]
			}
			if nxt == n*n { // 到达终点
				return p.step + 1
			}
			if !visited[nxt] {
				visited[nxt] = true
				queue = append(queue, pair{nxt, p.step + 1}) // 扩展新状态
			}
		}
	}
	return -1
}

func getRowCol(id, n int) (r, c int) {
	r, c = (id-1)/n, (id-1)%n
	if r%2 == 1 {
		c = n - 1 - c
	}
	r = n - 1 - r
	return r, c
}
