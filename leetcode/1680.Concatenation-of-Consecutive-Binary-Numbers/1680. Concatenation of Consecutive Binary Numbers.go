package leetcode

import (
	"math/bits"
)

// 解法一 模拟
func concatenatedBinary(n int) int {
	res, mod, shift := 0, 1000000007, 0
	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 {
			shift++
		}
		res = ((res << shift) + i) % mod
	}
	return res
}

// 解法二 位运算
func concatenatedBinary1(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res = (res<<bits.Len(uint(i)) | i) % (1e9 + 7)
	}
	return res
}
