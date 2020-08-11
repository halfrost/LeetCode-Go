package leetcode

func luckyNumbers(matrix [][]int) []int {
	t, r, res := make([]int, len(matrix[0])), make([]int, len(matrix[0])), []int{}
	for _, val := range matrix {
		m, k := val[0], 0
		for j := 0; j < len(matrix[0]); j++ {
			if val[j] < m {
				m = val[j]
				k = j
			}
			if t[j] < val[j] {
				t[j] = val[j]
			}
		}

		if t[k] == m {
			r[k] = m
		}
	}
	for k, v := range r {
		if v > 0 && v == t[k] {
			res = append(res, v)
		}
	}
	return res
}
