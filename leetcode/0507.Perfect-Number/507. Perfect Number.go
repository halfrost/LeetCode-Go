package leetcode

import "math"

// 方法一
func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum, bound := 1, int(math.Sqrt(float64(num)))+1
	for i := 2; i < bound; i++ {
		if num%i != 0 {
			continue
		}
		corrDiv := num / i
		sum += corrDiv + i
	}
	return sum == num
}

// 方法二 打表
func checkPerfectNumber_(num int) bool {
	return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
}
