package leetcode

// 解法一
func hasAlternatingBits(n int) bool {
	/*
	   n =         1 0 1 0 1 0 1 0
	   n >> 1      0 1 0 1 0 1 0 1
	   n ^ n>>1    1 1 1 1 1 1 1 1
	   n           1 1 1 1 1 1 1 1
	   n + 1     1 0 0 0 0 0 0 0 0
	   n & (n+1)   0 0 0 0 0 0 0 0
	*/
	n = n ^ (n >> 1)
	return (n & (n + 1)) == 0
}

// 解法二
func hasAlternatingBits1(n int) bool {
	last, current := 0, 0
	for n > 0 {
		last = n & 1
		n = n / 2
		current = n & 1
		if last == current {
			return false
		}
	}
	return true
}
