package leetcode

func minAddToMakeValid(S string) int {
	if len(S) == 0 {
		return 0
	}
	stack := make([]rune, 0)
	for _, v := range S {
		if v == '(' {
			stack = append(stack, v)
		} else if (v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack)
}
