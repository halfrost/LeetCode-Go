package leetcode

func findBestValue(arr []int, target int) int {
	low, high := 0, 100000
	for low < high {
		mid := low + (high-low)>>1
		if calculateSum(arr, mid) < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if high == 100000 {
		res := 0
		for _, num := range arr {
			if res < num {
				res = num
			}
		}
		return res
	}
	// 比较阈值线分别定在 left - 1 和 left 的时候与 target 的接近程度
	sum1, sum2 := calculateSum(arr, low-1), calculateSum(arr, low)
	if target-sum1 <= sum2-target {
		return low - 1
	}
	return low
}

func calculateSum(arr []int, mid int) int {
	sum := 0
	for _, num := range arr {
		sum += min(num, mid)
	}
	return sum
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
