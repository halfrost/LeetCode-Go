package leetcode

import (
	"math"
)

// 解法一
func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)
	for i := range res {
		res[i] = n
	}
	for i := 0; i < n; i++ {
		if s[i] == c {
			res[i] = 0
		} else if i > 0 {
			res[i] = res[i-1] + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && res[i+1]+1 < res[i] {
			res[i] = res[i+1] + 1
		}
	}
	return res
}

// 解法二
func shortestToChar1(s string, c byte) []int {
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
