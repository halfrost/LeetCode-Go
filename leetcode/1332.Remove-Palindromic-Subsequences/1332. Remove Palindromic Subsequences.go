package leetcode

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return 2
		}
	}
	return 1
}
