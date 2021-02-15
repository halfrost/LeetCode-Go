package leetcode

func minOperations(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') != i%2 {
			res++
		}
	}
	return min(res, len(s)-res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
