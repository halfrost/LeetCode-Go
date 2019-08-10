package leetcode

import "math"

func powerfulIntegers(x int, y int, bound int) []int {
	if x == 1 && y == 1 {
		if bound < 2 {
			return []int{}
		}
		return []int{2}
	}
	if x > y {
		x, y = y, x
	}
	visit, result := make(map[int]bool), make([]int, 0)
	for i := 0; ; i++ {
		found := false
		for j := 0; pow(x, i)+pow(y, j) <= bound; j++ {
			v := pow(x, i) + pow(y, j)
			if !visit[v] {
				found = true
				visit[v] = true
				result = append(result, v)
			}
		}
		if !found {
			break
		}
	}
	return result
}

func pow(x, i int) int {
	return int(math.Pow(float64(x), float64(i)))
}
