package leetcode

func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	if A[0] < A[1] {
		return inc(A[1:])
	}
	if A[0] > A[1] {
		return dec(A[1:])
	}
	return inc(A[1:]) || dec(A[1:])
}

func inc(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func dec(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			return false
		}
	}
	return true
}
