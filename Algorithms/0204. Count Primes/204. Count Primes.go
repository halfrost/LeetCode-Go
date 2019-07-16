package leetcode

func countPrimes(n int) int {
	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * i; j < n; j = j + i {
			isNotPrime[j] = true
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			count++
		}
	}
	return count
}
