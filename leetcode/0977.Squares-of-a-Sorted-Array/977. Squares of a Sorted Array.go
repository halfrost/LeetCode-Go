package leetcode

import "sort"

// 解法一
func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	for i, k, j := 0, len(A)-1, len(ans)-1; i <= j; k-- {
		if A[i]*A[i] > A[j]*A[j] {
			ans[k] = A[i] * A[i]
			i++
		} else {
			ans[k] = A[j] * A[j]
			j--
		}
	}
	return ans
}

// 解法二
func sortedSquares1(A []int) []int {
	for i, value := range A {
		A[i] = value * value
	}
	sort.Ints(A)
	return A
}
