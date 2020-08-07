package leetcode

func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for ; n > 0; n /= 10 {
		sum += n % 10
		product *= n % 10
	}
	return product - sum
}
