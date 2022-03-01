package leetcode

import "unicode"

// 解法一 分治，时间复杂度 O(n)
func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}

	chars := map[rune]int{}
	for _, r := range s {
		chars[r]++
	}

	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		_, u := chars[unicode.ToUpper(r)]
		_, l := chars[unicode.ToLower(r)]
		if u && l {
			continue
		}
		left := longestNiceSubstring(s[:i])
		right := longestNiceSubstring(s[i+1:])
		if len(left) >= len(right) {
			return left
		} else {
			return right
		}
	}
	return s
}

// 解法二 用二进制表示状态
func longestNiceSubstring1(s string) (ans string) {
	for i := range s {
		lower, upper := 0, 0
		for j := i; j < len(s); j++ {
			if unicode.IsLower(rune(s[j])) {
				lower |= 1 << (s[j] - 'a')
			} else {
				upper |= 1 << (s[j] - 'A')
			}
			if lower == upper && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return
}

// 解法三 暴力枚举，时间复杂度 O(n^2)
func longestNiceSubstring2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		m := map[byte]int{}
		m[s[i]]++
		for j := i + 1; j < len(s); j++ {
			m[s[j]]++
			if checkNiceString(m) && (j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func checkNiceString(m map[byte]int) bool {
	for k := range m {
		if k >= 97 && k <= 122 {
			if _, ok := m[k-32]; !ok {
				return false
			}
		}
		if k >= 65 && k <= 90 {
			if _, ok := m[k+32]; !ok {
				return false
			}
		}
	}
	return true
}
