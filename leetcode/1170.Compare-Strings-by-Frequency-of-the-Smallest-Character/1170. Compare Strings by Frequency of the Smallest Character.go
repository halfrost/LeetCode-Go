package leetcode

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	ws, res := make([]int, len(words)), make([]int, len(queries))
	for i, w := range words {
		ws[i] = countFunc(w)
	}
	sort.Ints(ws)
	for i, q := range queries {
		fq := countFunc(q)
		res[i] = len(words) - sort.Search(len(words), func(i int) bool { return fq < ws[i] })
	}
	return res
}

func countFunc(s string) int {
	count, i := [26]int{}, 0
	for _, b := range s {
		count[b-'a']++
	}
	for count[i] == 0 {
		i++
	}
	return count[i]
}
