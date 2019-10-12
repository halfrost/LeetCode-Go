package leetcode

func isPerfectSquare(num int) bool {
	low, high := 1, num
	for low <= high {
		mid := low + (high-low)>>1
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
