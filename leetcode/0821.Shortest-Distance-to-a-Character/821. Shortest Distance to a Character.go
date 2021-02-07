package leetcode

import (
	"math"
)

func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			res[i] = 0
		} else {
			left, right := math.MaxInt32, math.MaxInt32
			for j := i + 1; j < len(s); j++ {
				if s[j] == c {
					right = j - i
					break
				}
			}
			for k := i - 1; k >= 0; k-- {
				if s[k] == c {
					left = i - k
					break
				}
			}
			res[i] = min(left, right)
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
