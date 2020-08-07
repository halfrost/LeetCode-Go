package leetcode

func numEquivDominoPairs(dominoes [][]int) int {
	if dominoes == nil || len(dominoes) == 0 {
		return 0
	}
	result, buckets := 0, [100]int{}
	for _, dominoe := range dominoes {
		key, rotatedKey := dominoe[0]*10+dominoe[1], dominoe[1]*10+dominoe[0]
		if dominoe[0] != dominoe[1] {
			if buckets[rotatedKey] > 0 {
				result += buckets[rotatedKey]
			}
		}
		if buckets[key] > 0 {
			result += buckets[key]
			buckets[key]++
		} else {
			buckets[key]++
		}
	}
	return result
}
