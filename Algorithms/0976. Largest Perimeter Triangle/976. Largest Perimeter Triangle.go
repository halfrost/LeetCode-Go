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

func quickSort164(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition164(a, lo, hi)
	quickSort164(a, lo, p-1)
	quickSort164(a, p+1, hi)
}

func partition164(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
