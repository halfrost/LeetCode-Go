package leetcode

import "math"

func judgeSquareSum(c int) bool {
	low, high := 0, int(math.Sqrt(float64(c)))
	for low <= high {
		if low*low+high*high < c {
			low++
		} else if low*low+high*high > c {
			high--
		} else {
			return true
		}
	}
	return false
}
