package leetcode

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}
	res, dp := 0, 0
	for i := 1; i < len(A)-1; i++ {
		if A[i+1]-A[i] == A[i]-A[i-1] {
			dp++
			res += dp
		} else {
			dp = 0
		}
	}
	return res
}
