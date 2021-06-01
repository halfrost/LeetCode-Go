package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		v, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, v)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
