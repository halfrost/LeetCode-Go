package leetcode

func countTriplets(arr []int) int {
	prefix, num, count, total := make([]int, len(arr)), 0, 0, 0
	for i, v := range arr {
		num ^= v
		prefix[i] = num
	}
	for i := 0; i < len(prefix)-1; i++ {
		for k := i + 1; k < len(prefix); k++ {
			total = prefix[k]
			if i > 0 {
				total ^= prefix[i-1]
			}
			if total == 0 {
				count += k - i
			}
		}
	}
	return count
}
