package leetcode

func addDigits(num int) int {
	for num > 9 {
		cur := 0
		for num != 0 {
			cur += num % 10
			num /= 10
		}
		num = cur
	}
	return num
}
