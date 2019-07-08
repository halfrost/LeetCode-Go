package leetcode

// 解法一
func rangeBitwiseAnd1(m int, n int) int {
	if m == 0 {
		return 0
	}
	moved := 0
	for m != n {
		m >>= 1
		n >>= 1
		moved++
	}
	return m << uint32(moved)
}

// 解法二 Brian Kernighan's algorithm
func rangeBitwiseAnd(m int, n int) int {
	for n > m {
		n &= (n - 1) // 清除最低位的 1
	}
	return n
}
