package leetcode

// 解法一
func numberOfMatches(n int) int {
	return n - 1
}

// 解法二 模拟
func numberOfMatches1(n int) int {
	sum := 0
	for n != 1 {
		if n&1 == 0 {
			sum += n / 2
			n = n / 2
		} else {
			sum += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return sum
}
