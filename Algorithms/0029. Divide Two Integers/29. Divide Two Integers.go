package leetcode

import (
	"math"
)

// 解法一 递归版的二分搜索
func divide(dividend int, divisor int) int {
	sign, res := -1, 0
	// low, high := 0, abs(dividend)
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}
	if dividend > math.MaxInt32 {
		dividend = math.MaxInt32
	}
	// 如果把递归改成非递归，可以改成下面这段代码
	// for low <= high {
	// 	quotient := low + (high-low)>>1
	// 	if ((quotient+1)*abs(divisor) > abs(dividend) && quotient*abs(divisor) <= abs(dividend)) || ((quotient+1)*abs(divisor) >= abs(dividend) && quotient*abs(divisor) < abs(dividend)) {
	// 		if (quotient+1)*abs(divisor) == abs(dividend) {
	// 			res = quotient + 1
	// 			break
	// 		}
	// 		res = quotient
	// 		break
	// 	}
	// 	if (quotient+1)*abs(divisor) > abs(dividend) && quotient*abs(divisor) > abs(dividend) {
	// 		high = quotient - 1
	// 	}
	// 	if (quotient+1)*abs(divisor) < abs(dividend) && quotient*abs(divisor) < abs(dividend) {
	// 		low = quotient + 1
	// 	}
	// }
	res = binarySearchQuotient(0, abs(dividend), abs(divisor), abs(dividend))
	if res > math.MaxInt32 {
		return sign * math.MaxInt32
	}
	if res < math.MinInt32 {
		return sign * math.MinInt32
	}
	return sign * res
}

func binarySearchQuotient(low, high, val, dividend int) int {
	quotient := low + (high-low)>>1
	if ((quotient+1)*val > dividend && quotient*val <= dividend) || ((quotient+1)*val >= dividend && quotient*val < dividend) {
		if (quotient+1)*val == dividend {
			return quotient + 1
		}
		return quotient
	}
	if (quotient+1)*val > dividend && quotient*val > dividend {
		return binarySearchQuotient(low, quotient-1, val, dividend)
	}
	if (quotient+1)*val < dividend && quotient*val < dividend {
		return binarySearchQuotient(quotient+1, high, val, dividend)
	}
	return 0
}

// 解法二 非递归版的二分搜索
func divide1(divided int, divisor int) int {
	if divided == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	result := 0
	sign := -1
	if divided > 0 && divisor > 0 || divided < 0 && divisor < 0 {
		sign = 1
	}
	dvd, dvs := abs(divided), abs(divisor)
	for dvd >= dvs {
		temp := dvs
		m := 1
		for temp<<1 <= dvd {
			temp <<= 1
			m <<= 1
		}
		dvd -= temp
		result += m
	}
	return sign * result
}
