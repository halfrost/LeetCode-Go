package leetcode

func shipWithinDays(weights []int, D int) int {
	maxNum, sum := 0, 0
	for _, num := range weights {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	if D == 1 {
		return sum
	}
	low, high := maxNum, sum
	for low < high {
		mid := low + (high-low)>>1
		if calSum(mid, D, weights) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
