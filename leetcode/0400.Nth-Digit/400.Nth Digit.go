package leetcode

import "math"

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	bits := 1
	for n > 9*int(math.Pow10(bits-1))*bits {
		n -= 9 * int(math.Pow10(bits-1)) * bits
		bits++
	}
	idx := n - 1
	start := int(math.Pow10(bits - 1))
	num := start + idx/bits
	digitIdx := idx % bits
	return num / int(math.Pow10(bits-digitIdx-1)) % 10
}
