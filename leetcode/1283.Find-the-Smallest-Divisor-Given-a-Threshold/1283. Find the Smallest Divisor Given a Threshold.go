package leetcode

func smallestDivisor(nums []int, threshold int) int {
	low, high := 1, 1000000
	for low < high {
		mid := low + (high-low)>>1
		if calDivisor(nums, mid, threshold) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func calDivisor(nums []int, mid, threshold int) bool {
	sum := 0
	for i := range nums {
		if nums[i]%mid != 0 {
			sum += nums[i]/mid + 1
		} else {
			sum += nums[i] / mid
		}
	}
	if sum <= threshold {
		return true
	}
	return false
}
