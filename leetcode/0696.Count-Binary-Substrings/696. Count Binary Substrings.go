package leetcode

func countBinarySubstrings(s string) int {
	last, res := 0, 0
	for i := 0; i < len(s); {
		c, count := s[i], 1
		for i++; i < len(s) && s[i] == c; i++ {
			count++
		}
		res += min(count, last)
		last = count
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
