package leetcode

func nthUglyNumber(n int, a int, b int, c int) int {
	low, high := int64(0), int64(2*1e9)
	for low < high {
		mid := low + (high-low)>>1
		if calNthCount(mid, int64(a), int64(b), int64(c)) < int64(n) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return int(low)
}

func calNthCount(num, a, b, c int64) int64 {
	ab, bc, ac := a*b/gcd(a, b), b*c/gcd(b, c), a*c/gcd(a, c)
	abc := a * bc / gcd(a, bc)
	return num/a + num/b + num/c - num/ab - num/bc - num/ac + num/abc
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
