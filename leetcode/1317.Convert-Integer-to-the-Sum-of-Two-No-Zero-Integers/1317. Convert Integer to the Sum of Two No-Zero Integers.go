package leetcode

func getNoZeroIntegers(n int) []int {
	noZeroPair := []int{}
	for i := 1; i <= n/2; i++ {
		if isNoZero(i) && isNoZero(n-i) {
			noZeroPair = append(noZeroPair, []int{i, n - i}...)
			break
		}
	}
	return noZeroPair
}

func isNoZero(n int) bool {
	for n != 0 {
		if n%10 == 0 {
			return false
		}
		n /= 10
	}
	return true
}
