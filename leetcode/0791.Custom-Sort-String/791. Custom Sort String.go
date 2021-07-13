package leetcode

import "sort"

func customSortString(order string, str string) string {
	magic := map[byte]int{}
	for i := range order {
		magic[order[i]] = i - 30
	}
	byteSlice := []byte(str)
	sort.Slice(byteSlice, func(i, j int) bool {
		return magic[byteSlice[i]] < magic[byteSlice[j]]
	})
	return string(byteSlice)
}
