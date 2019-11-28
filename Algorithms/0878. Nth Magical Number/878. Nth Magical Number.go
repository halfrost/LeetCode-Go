package leetcode

func nthMagicalNumber(N int, A int, B int) int {
	low, high := int64(0), int64(1*1e14)
	for low < high {
		mid := low + (high-low)>>1
		if calNthMagicalCount(mid, int64(A), int64(B)) < int64(N) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return int(low) % 1000000007
}

func calNthMagicalCount(num, a, b int64) int64 {
	ab := a * b / gcd(a, b)
	return num/a + num/b - num/ab
}
