package leetcode

func splitArray(nums []int, m int) int {
	maxNum, sum := 0, 0
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	if m == 1 {
		return sum
	}
	low, high := maxNum, sum
	for low < high {
		mid := low + (high-low)>>1
		if calSum(mid, m, nums) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func calSum(mid, m int, nums []int) bool {
	sum, count := 0, 0
	for _, v := range nums {
		sum += v
		if sum > mid {
			sum = v
			count++
			// 分成 m 块，只需要插桩 m -1 个
			if count > m-1 {
				return false
			}
		}
	}
	return true
}
