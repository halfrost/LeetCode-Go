package leetcode

func isValid1003(S string) bool {
	if len(S) < 3 {
		return false
	}
	stack := []byte{}
	for i := 0; i < len(S); i++ {
		if S[i] == 'a' {
			stack = append(stack, S[i])
		} else if S[i] == 'b' {
			if len(stack) > 0 && stack[len(stack)-1] == 'a' {
				stack = append(stack, S[i])
			} else {
				return false
			}
		} else {
			if len(stack) > 1 && stack[len(stack)-1] == 'b' && stack[len(stack)-2] == 'a' {
				stack = stack[:len(stack)-2]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
