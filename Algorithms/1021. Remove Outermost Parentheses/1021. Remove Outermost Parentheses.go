package leetcode

// 解法一
func removeOuterParentheses(S string) string {
	now, current, ans := 0, "", ""
	for _, char := range S {
		if string(char) == "(" {
			now++
		} else if string(char) == ")" {
			now--
		}
		current += string(char)
		if now == 0 {
			ans += current[1 : len(current)-1]
			current = ""
		}
	}
	return ans
}

// 解法二
func removeOuterParentheses1(S string) string {
	stack, res, counter := []byte{}, "", 0
	for i := 0; i < len(S); i++ {
		if counter == 0 && len(stack) == 1 && S[i] == ')' {
			stack = stack[1:]
			continue
		}
		if len(stack) == 0 && S[i] == '(' {
			stack = append(stack, S[i])
			continue
		}
		if len(stack) > 0 {
			switch S[i] {
			case '(':
				{
					counter++
					res += "("
				}
			case ')':
				{
					counter--
					res += ")"
				}
			}
		}
	}
	return res
}
