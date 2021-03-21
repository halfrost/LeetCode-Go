package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	stack, j, N := []int{}, 0, len(pushed)
	for _, x := range pushed {
		stack = append(stack, x)
		for len(stack) != 0 && j < N && stack[len(stack)-1] == popped[j] {
			stack = stack[0 : len(stack)-1]
			j++
		}
	}
	return j == N
}
