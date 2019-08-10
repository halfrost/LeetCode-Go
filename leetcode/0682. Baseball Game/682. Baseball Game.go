package leetcode

import "strconv"

func calPoints(ops []string) int {
	stack := make([]int, len(ops))
	top := 0

	for i := 0; i < len(ops); i++ {
		op := ops[i]
		switch op {
		case "+":
			last1 := stack[top-1]
			last2 := stack[top-2]
			stack[top] = last1 + last2
			top++
		case "D":
			last1 := stack[top-1]
			stack[top] = last1 * 2
			top++
		case "C":
			top--
		default:
			stack[top], _ = strconv.Atoi(op)
			top++
		}
	}

	points := 0
	for i := 0; i < top; i++ {
		points += stack[i]
	}
	return points
}
