package leetcode

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	num, _ := strconv.ParseUint(n, 10, 64)
	for bit := uint64(math.Log2(float64(num))); bit >= 1; bit-- {
		low, high := uint64(2), uint64(math.Pow(float64(num), 1.0/float64(bit)))
		for low < high {
			mid := uint64(low + (high-low)>>1)
			sum := findBase(mid, bit)
			if sum == num {
				return strconv.FormatUint(mid, 10)
			} else if sum > num {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return strconv.FormatUint(num-1, 10)
}

// 计算 k^m + k^(m-1) + ... + k + 1
func findBase(mid, bit uint64) uint64 {
	sum, base := uint64(1), uint64(1)
	for i := uint64(1); i <= bit; i++ {
		base *= mid
		sum += base
	}
	return sum
}
