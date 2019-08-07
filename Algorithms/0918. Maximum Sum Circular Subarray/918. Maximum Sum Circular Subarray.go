package leetcode

import "math"

func maxSubarraySumCircular(A []int) int {
	n, sum := len(A), 0
	for _, v := range A {
		sum += v
	}
	kad := kadane(A)
	for i := 0; i < n; i++ {
		A[i] = -A[i]
	}
	negativeMax := kadane(A)
	if sum+negativeMax <= 0 {
		return kad
	}
	return max(kad, sum+negativeMax)
}

func kadane(a []int) int {
	n, MaxEndingHere, maxSoFar := len(a), a[0], math.MinInt32
	for i := 1; i < n; i++ {
		MaxEndingHere = max(a[i], MaxEndingHere+a[i])
		maxSoFar = max(MaxEndingHere, maxSoFar)
	}
	return maxSoFar
}
