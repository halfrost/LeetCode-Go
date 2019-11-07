package leetcode

import "math"

func findKthNumber(m int, n int, k int) int {
	low, high := 1, m*n
	for low < high {
		mid := low + (high-low)>>1
		if counterKthNum(m, n, mid) >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func counterKthNum(m, n, mid int) int {
	count := 0
	for i := 1; i <= m; i++ {
		count += int(math.Min(math.Floor(float64(mid)/float64(i)), float64(n)))
	}
	return count
}
