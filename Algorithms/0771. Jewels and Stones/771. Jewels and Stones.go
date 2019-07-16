package leetcode

import "strings"

// 解法一
func numJewelsInStones(J string, S string) int {
	count := 0
	for i := range S {
		if strings.Contains(J, string(S[i])) {
			count++
		}
	}
	return count
}

// 解法二
func numJewelsInStones1(J string, S string) int {
	cache, result := make(map[rune]bool), 0
	for _, r := range J {
		cache[r] = true
	}
	for _, r := range S {
		if _, ok := cache[r]; ok {
			result++
		}
	}
	return result
}
