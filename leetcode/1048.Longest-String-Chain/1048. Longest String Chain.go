package leetcode

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	poss, res := make([]int, 16+2), 0
	for i, w := range words {
		if poss[len(w)] == 0 {
			poss[len(w)] = i
		}
	}
	dp := make([]int, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		dp[i] = 1
		for j := poss[len(words[i])+1]; j < len(words) && len(words[j]) == len(words[i])+1; j++ {
			if isPredecessor(words[j], words[i]) {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPredecessor(long, short string) bool {
	i, j := 0, 0
	wasMismatch := false
	for j < len(short) {
		if long[i] != short[j] {
			if wasMismatch {
				return false
			}
			wasMismatch = true
			i++
			continue
		}
		i++
		j++
	}
	return true
}
