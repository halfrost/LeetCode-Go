package leetcode

func isIdealPermutation(A []int) bool {
	for i := range A {
		if abs(A[i]-i) > 1 {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
