package leetcode

func hIndex275(citations []int) int {
	low, high := 0, len(citations)-1
	for low <= high {
		mid := low + (high-low)>>1
		if len(citations)-mid > citations[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return len(citations) - low
}
