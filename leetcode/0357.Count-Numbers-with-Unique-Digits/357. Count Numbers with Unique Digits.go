package leetcode

// 暴力打表法
func countNumbersWithUniqueDigits1(n int) int {
	res := []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851, 5611771, 8877691}
	if n >= 10 {
		return res[10]
	}
	return res[n]
}

// 打表方法
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	res, uniqueDigits, availableNumber := 10, 9, 9
	for n > 1 && availableNumber > 0 {
		uniqueDigits = uniqueDigits * availableNumber
		res += uniqueDigits
		availableNumber--
		n--
	}
	return res
}
