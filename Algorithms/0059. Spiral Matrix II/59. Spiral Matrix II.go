package leetcode

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{[]int{1}}
	}
	res, visit, round, x, y, spDir := make([][]int, n), make([][]int, n), 0, 0, 0, [][]int{
		[]int{0, 1},  // 朝右
		[]int{1, 0},  // 朝下
		[]int{0, -1}, // 朝左
		[]int{-1, 0}, // 朝上
	}
	for i := 0; i < n; i++ {
		visit[i] = make([]int, n)
		res[i] = make([]int, n)
	}
	visit[x][y] = 1
	res[x][y] = 1
	for i := 0; i < n*n; i++ {
		x += spDir[round%4][0]
		y += spDir[round%4][1]
		if (x == 0 && y == n-1) || (x == n-1 && y == n-1) || (y == 0 && x == n-1) {
			round++
		}
		if x > n-1 || y > n-1 || x < 0 || y < 0 {
			return res
		}
		if visit[x][y] == 0 {
			visit[x][y] = 1
			res[x][y] = i + 2
		}
		switch round % 4 {
		case 0:
			if y+1 <= n-1 && visit[x][y+1] == 1 {
				round++
				continue
			}
		case 1:
			if x+1 <= n-1 && visit[x+1][y] == 1 {
				round++
				continue
			}
		case 2:
			if y-1 >= 0 && visit[x][y-1] == 1 {
				round++
				continue
			}
		case 3:
			if x-1 >= 0 && visit[x-1][y] == 1 {
				round++
				continue
			}
		}
	}
	return res
}
