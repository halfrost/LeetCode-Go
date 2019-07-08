package leetcode

// 解法一 数论
func isPowerOfThree(n int) bool {
	// 1162261467 is 3^19,  3^20 is bigger than int
	return n > 0 && (1162261467%n == 0)
}

// 解法二 打表法
func isPowerOfThree1(n int) bool {
	// 1162261467 is 3^19,  3^20 is bigger than int
	allPowerOfThreeMap := map[int]int{1: 1, 3: 3, 9: 9, 27: 27, 81: 81, 243: 243, 729: 729, 2187: 2187, 6561: 6561, 19683: 19683, 59049: 59049, 177147: 177147, 531441: 531441, 1594323: 1594323, 4782969: 4782969, 14348907: 14348907, 43046721: 43046721, 129140163: 129140163, 387420489: 387420489, 1162261467: 1162261467}
	_, ok := allPowerOfThreeMap[n]
	return ok
}

// 解法三 循环
func isPowerOfThree2(num int) bool {
	for num >= 3 {
		if num%3 == 0 {
			num = num / 3
		} else {
			return false
		}
	}
	return num == 1
}
