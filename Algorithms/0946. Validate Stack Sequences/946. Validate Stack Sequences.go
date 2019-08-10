package leetcode

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack, j, N := []int{}, 0, len(pushed)
	for _, x := range pushed {
		stack = append(stack, x)
		fmt.Printf("stack = %v j = %v\n", stack, j)
		for len(stack) != 0 && j < N && stack[len(stack)-1] == popped[j] {
			stack = stack[0 : len(stack)-1]
			j++
		}
		fmt.Printf("*****stack = %v j = %v\n", stack, j)
	}
	return j == N
}
