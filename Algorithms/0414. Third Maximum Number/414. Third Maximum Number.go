package leetcode

import (
	"math"
)

func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v > a {
			c = b
			b = a
			a = v
		} else if v < a && v > b {
			c = b
			b = v
		} else if v < b && v > c {
			c = v
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
