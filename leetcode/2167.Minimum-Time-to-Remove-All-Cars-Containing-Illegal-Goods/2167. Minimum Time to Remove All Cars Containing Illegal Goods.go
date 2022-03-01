package leetcode

import "runtime/debug"

// 解法一 DP
func minimumTime(s string) int {
	suffixSum, prefixSum, res := make([]int, len(s)+1), make([]int, len(s)+1), 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			suffixSum[i] = suffixSum[i+1]
		} else {
			suffixSum[i] = min(suffixSum[i+1]+2, len(s)-i)
		}
	}
	res = suffixSum[0]
	if s[0] == '1' {
		prefixSum[0] = 1
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			prefixSum[i] = prefixSum[i-1]
		} else {
			prefixSum[i] = min(prefixSum[i-1]+2, i+1)
		}
		res = min(res, prefixSum[i]+suffixSum[i+1])
	}
	return res
}

func init() { debug.SetGCPercent(-1) }

// 解法二 小幅优化时间和空间复杂度
func minimumTime1(s string) int {
	res, count := len(s), 0
	for i := 0; i < len(s); i++ {
		count = min(count+int(s[i]-'0')*2, i+1)
		res = min(res, count+len(s)-i-1)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
