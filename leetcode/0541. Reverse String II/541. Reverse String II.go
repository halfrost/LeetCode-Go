package leetcode

func reverseStr(s string, k int) string {
	if k > len(s) {
		k = len(s)
	}
	for i := 0; i < len(s); i = i + 2*k {
		if len(s)-i >= k {
			ss := revers(s[i : i+k])
			s = s[:i] + ss + s[i+k:]
		} else {
			ss := revers(s[i:])
			s = s[:i] + ss
		}
	}
	return s
}
