package leetcode

func scoreOfParentheses(S string) int {
	res, stack, top, temp := 0, []int{}, -1, 0
	for _, s := range S {
		if s == '(' {
			stack = append(stack, -1)
			top++
		} else {
			temp = 0
			for stack[top] != -1 {
				temp += stack[top]
				stack = stack[:len(stack)-1]
				top--
			}
			stack = stack[:len(stack)-1]
			top--
			if temp == 0 {
				stack = append(stack, 1)
				top++
			} else {
				stack = append(stack, temp*2)
				top++
			}
		}
	}
	for len(stack) != 0 {
		res += stack[top]
		stack = stack[:len(stack)-1]
		top--
	}
	return res
}
