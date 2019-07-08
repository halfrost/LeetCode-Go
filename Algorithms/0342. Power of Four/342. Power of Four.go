package leetcode

// 解法一 数论
func isPowerOfFour(num int) bool {
	return num > 0 && (num&(num-1)) == 0 && (num-1)%3 == 0
}

// 解法二 循环
func isPowerOfFour1(num int) bool {
	for num >= 4 {
		if num%4 == 0 {
			num = num / 4
		} else {
			return false
		}
	}
	return num == 1
}
