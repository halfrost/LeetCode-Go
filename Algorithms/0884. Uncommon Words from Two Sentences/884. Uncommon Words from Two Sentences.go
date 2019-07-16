package leetcode

import "strings"

func uncommonFromSentences(A string, B string) []string {
	m, res := map[string]int{}, []string{}
	for _, s := range []string{A, B} {
		for _, word := range strings.Split(s, " ") {
			m[word]++
		}
	}
	for key := range m {
		if m[key] == 1 {
			res = append(res, key)
		}
	}
	return res
}
