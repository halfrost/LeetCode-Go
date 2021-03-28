package leetcode

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += countPalindrome(s, i, i)
		res += countPalindrome(s, i, i+1)
	}
	return res
}

func countPalindrome(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
		res++
	}
	return res
}
