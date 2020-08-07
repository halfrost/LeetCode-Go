package leetcode

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	res, round, spDir := [][]int{}, 0, [][]int{
		[]int{0, 1},  // 朝右
		[]int{1, 0},  // 朝下
		[]int{0, -1}, // 朝左
		[]int{-1, 0}, // 朝上
	}
	res = append(res, []int{r0, c0})
	for i := 0; len(res) < R*C; i++ {
		for j := 0; j < i/2+1; j++ {
			r0 += spDir[round%4][0]
			c0 += spDir[round%4][1]
			if 0 <= r0 && r0 < R && 0 <= c0 && c0 < C {
				res = append(res, []int{r0, c0})
			}
		}
		round++
	}
	return res
}
