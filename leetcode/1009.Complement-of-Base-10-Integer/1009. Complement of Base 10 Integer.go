package leetcode

func bitwiseComplement(n int) int {
	mask := 1
	for mask < n {
		mask = (mask << 1) + 1
	}
	return mask ^ n
}
