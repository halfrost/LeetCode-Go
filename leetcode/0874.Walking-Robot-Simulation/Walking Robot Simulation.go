package leetcode

func robotSim(commands []int, obstacles [][]int) int {
	m := make(map[[2]int]struct{})
	for _, v := range obstacles {
		if len(v) != 0 {
			m[[2]int{v[0], v[1]}] = struct{}{}
		}
	}
	directX := []int{0, 1, 0, -1}
	directY := []int{1, 0, -1, 0}
	direct, x, y := 0, 0, 0
	result := 0
	for _, c := range commands {
		if c == -2 {
			direct = (direct + 3) % 4
			continue
		}
		if c == -1 {
			direct = (direct + 1) % 4
			continue
		}
		for ; c > 0; c-- {
			nextX := x + directX[direct]
			nextY := y + directY[direct]
			if _, ok := m[[2]int{nextX, nextY}]; ok {
				break
			}
			tmpResult := nextX*nextX + nextY*nextY
			if tmpResult > result {
				result = tmpResult
			}
			x = nextX
			y = nextY
		}
	}
	return result
}
