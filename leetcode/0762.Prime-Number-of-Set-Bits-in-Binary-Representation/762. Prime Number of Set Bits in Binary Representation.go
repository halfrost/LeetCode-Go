package leetcode

import "math/bits"

func countPrimeSetBits(L int, R int) int {
	counter := 0
	for i := L; i <= R; i++ {
		if isPrime(bits.OnesCount(uint(i))) {
			counter++
		}
	}
	return counter
}

func isPrime(x int) bool {
	return x == 2 || x == 3 || x == 5 || x == 7 || x == 11 || x == 13 || x == 17 || x == 19
}
