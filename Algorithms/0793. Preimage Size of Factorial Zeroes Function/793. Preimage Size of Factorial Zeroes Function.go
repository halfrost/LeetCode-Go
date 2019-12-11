package leetcode

// 解法一 二分搜索
func preimageSizeFZF(K int) int {
	low, high := 0, 5*K
	for low <= high {
		mid := low + (high-low)>>1
		k := trailingZeroes(mid)
		if k == K {
			return 5
		} else if k > K {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}

// 解法二 数学方法
func preimageSizeFZF1(K int) int {
	base := 0
	for base < K {
		base = base*5 + 1
	}
	for K > 0 {
		base = (base - 1) / 5
		if K/base == 5 {
			return 0
		}
		K %= base
	}
	return 5
}
