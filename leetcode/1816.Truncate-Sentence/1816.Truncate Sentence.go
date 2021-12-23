package leetcode

func truncateSentence(s string, k int) string {
	end := 0
	for i := range s {
		if k > 0 && s[i] == ' ' {
			k--
		}
		if k == 0 {
			end = i
			break
		}
	}
	if end == 0 {
		return s
	}
	return s[:end]
}
