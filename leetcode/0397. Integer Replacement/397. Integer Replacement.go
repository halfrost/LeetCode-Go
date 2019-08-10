package leetcode

func integerReplacement(n int) int {
	res := 0
	for n > 1 {
		if (n & 1) == 0 { // 判断是否是偶数
			n >>= 1
		} else if (n+1)%4 == 0 && n != 3 { // 末尾 2 位为 11
			n++
		} else { // 末尾 2 位为 01
			n--
		}
		res++
	}
	return res
}
