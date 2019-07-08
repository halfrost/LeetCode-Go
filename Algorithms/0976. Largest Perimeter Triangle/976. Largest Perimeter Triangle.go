package leetcode

func largestPerimeter(A []int) int {
	if len(A) < 3 {
		return 0
	}
	quickSort164(A, 0, len(A)-1)
	for i := len(A) - 1; i >= 2; i-- {
		if (A[i]+A[i-1] > A[i-2]) && (A[i]+A[i-2] > A[i-1]) && (A[i-2]+A[i-1] > A[i]) {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}
