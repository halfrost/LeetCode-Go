package leetcode

func minCut(s string) int {
	if s == "" {
		return 0
	}
	result := len(s)
	current := make([]string, 0, len(s))
	dfs132(s, 0, current, &result)
	return result
}

func dfs132(s string, idx int, cur []string, result *int) {
	start, end := idx, len(s)
	if start == end {
		*result = min(*result, len(cur)-1)
		return
	}
	for i := start; i < end; i++ {
		if isPal(s, start, i) {
			dfs132(s, i+1, append(cur, s[start:i+1]), result)
		}
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func isPal(str string, s, e int) bool {
	for s < e {
		if str[s] != str[e] {
			return false
		}
		s++
		e--
	}
	return true
}
