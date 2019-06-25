package leetcode

import (
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	mp := make(map[string]bool)
	var res string
	for _, word := range words {
		size := len(word)
		if size == 1 || mp[word[:size-1]] {
			if size > len(res) {
				res = word
			}
			mp[word] = true
		}
	}
	return res
}
