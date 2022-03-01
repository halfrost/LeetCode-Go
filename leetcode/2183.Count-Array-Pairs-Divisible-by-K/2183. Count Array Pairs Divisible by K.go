package leetcode

import "math"

func countPairs(nums []int, k int) int64 {
	n := int(math.Sqrt(float64(k)))
	gcds, res := make(map[int]int, n), 0
	for _, num := range nums {
		gcds[gcd(num, k)]++
	}

	for a, n1 := range gcds {
		for b, n2 := range gcds {
			if a > b || (a*b)%k != 0 {
				continue
			}
			if a != b {
				res += n1 * n2
			} else { // a == b
				res += n1 * (n1 - 1) / 2
			}
		}
	}
	return int64(res)
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
