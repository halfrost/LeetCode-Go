package leetcode

func constructArray(n int, k int) []int {
	res := []int{}
	for i := 0; i < n-k-1; i++ {
		res = append(res, i+1)
	}
	for i := n - k; i < n-k+(k+1)/2; i++ {
		res = append(res, i)
		res = append(res, 2*n-k-i)
	}
	if k%2 == 0 {
		res = append(res, n-k+(k+1)/2)
	}
	return res
}
