package leetcode

func diStringMatch(S string) []int {
	result, maxNum, minNum, index := make([]int, len(S)+1), len(S), 0, 0
	for _, ch := range S {
		if ch == 'I' {
			result[index] = minNum
			minNum++
		} else {
			result[index] = maxNum
			maxNum--
		}
		index++
	}
	result[index] = minNum
	return result
}
