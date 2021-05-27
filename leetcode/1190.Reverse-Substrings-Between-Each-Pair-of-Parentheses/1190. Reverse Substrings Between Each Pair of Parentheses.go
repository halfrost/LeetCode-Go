package leetcode

func reverseParentheses(s string) string {
	pair, stack := make([]int, len(s)), []int{}
	for i, b := range s {
		if b == '(' {
			stack = append(stack, i)
		} else if b == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i], pair[j] = j, i
		}
	}
	res := []byte{}
	for i, step := 0, 1; i < len(s); i += step {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			step = -step
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
