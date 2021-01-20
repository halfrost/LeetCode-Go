package leetcode

func largeGroupPositions(S string) [][]int {
	res, end := [][]int{}, 0
	for end < len(S) {
		start, str := end, S[end]
		for end < len(S) && S[end] == str {
			end++
		}
		if end-start >= 3 {
			res = append(res, []int{start, end - 1})
		}
	}
	return res
}
