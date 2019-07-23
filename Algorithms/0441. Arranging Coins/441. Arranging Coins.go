package leetcode

import "math"

// 解法一 数学公式
func arrangeCoins(n int) int {
	if n <= 0 {
		return 0
	}
	x := math.Sqrt(2*float64(n)+0.25) - 0.5
	return int(x)
}

// 解法二 模拟
func arrangeCoins1(n int) int {
	k := 1
	for n >= k {
		n -= k
		k++
	}
	return k - 1
}
