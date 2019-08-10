package leetcode

// 解法一
func findComplement(num int) int {
	xx := ^0 // ^0 = 1111111111111111111111
	for xx&num > 0 {
		xx <<= 1 // 构造出来的 xx = 1111111…000000，0 的个数就是 num 的长度
	}
	return ^xx ^ num // xx ^ num，结果是前面的 0 全是 1 的num，再取反即是答案
}

// 解法二
func findComplement1(num int) int {
	temp := 1
	for temp <= num {
		temp <<= 1 // 构造出来的 temp = 00000……10000，末尾 0 的个数是 num 的长度
	}
	return (temp - 1) ^ num // temp - 1 即是前面都是 0，num 长度的末尾都是 1 的数，再异或 num 即是最终结果
}
