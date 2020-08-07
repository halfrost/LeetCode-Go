package leetcode

import "sort"

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func numPrimeArrangements(n int) int {
	primeCount := sort.Search(25, func(i int) bool { return primes[i] > n })
	return factorial(primeCount) * factorial(n-primeCount) % 1000000007
}

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1) % 1000000007
}
