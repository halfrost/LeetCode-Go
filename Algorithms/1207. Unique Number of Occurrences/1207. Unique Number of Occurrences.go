package leetcode

func uniqueOccurrences(arr []int) bool {
	freq, m := map[int]int{}, map[int]bool{}
	for _, v := range arr {
		freq[v]++
	}
	for _, v := range freq {
		if _, ok := m[v]; !ok {
			m[v] = true
		} else {
			return false
		}
	}
	return true
}
