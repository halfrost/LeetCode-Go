package leetcode

import "sort"

func advantageCount1(A []int, B []int) []int {
	n := len(A)
	sort.Ints(A)
	sortedB := make([]int, n)
	for i := range sortedB {
		sortedB[i] = i
	}
	sort.Slice(sortedB, func(i, j int) bool {
		return B[sortedB[i]] < B[sortedB[j]]
	})
	useless, i, res := make([]int, 0), 0, make([]int, n)
	for _, index := range sortedB {
		b := B[index]
		for i < n && A[i] <= b {
			useless = append(useless, A[i])
			i++
		}
		if i < n {
			res[index] = A[i]
			i++
		} else {
			res[index] = useless[0]
			useless = useless[1:]
		}
	}
	return res
}
