package leetcode

import "strconv"

func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}
	res := ""
	for N != 0 {
		remainder := N % (-2)
		N = N / (-2)
		if remainder < 0 {
			remainder += 2
			N++
		}
		res = strconv.Itoa(remainder) + res
	}
	return res
}
