package leetcode

func minKBitFlips(A []int, K int) int {
	flippedTime, count := 0, 0
	for i := 0; i < len(A); i++ {
		if i >= K && A[i-K] == 2 {
			flippedTime--
		}
		// 下面这个判断包含了两种情况：
		// 如果 flippedTime 是奇数，且 A[i] == 1 就需要翻转
		// 如果 flippedTime 是偶数，且 A[i] == 0 就需要翻转
		if flippedTime%2 == A[i] {
			if i+K > len(A) {
				return -1
			}
			A[i] = 2
			flippedTime++
			count++
		}
	}
	return count
}
