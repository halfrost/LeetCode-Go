package leetcode

func maxDepth(s string) int {
	res, cur := 0, 0
	for _, c := range s {
		if c == '(' {
			cur++
			res = max(res, cur)
		} else if c == ')' {
			cur--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
