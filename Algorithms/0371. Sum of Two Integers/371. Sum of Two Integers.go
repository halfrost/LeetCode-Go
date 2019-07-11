package leetcode

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	// (a & b)<<1 计算的是进位
	// a ^ b 计算的是不带进位的加法
	return getSum((a&b)<<1, a^b)
}
