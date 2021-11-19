package leetcode

import "sort"

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) <= len(strs[j])
	})
	minLen := len(strs[0])
	if minLen == 0 {
		return ""
	}
	var commonPrefix []byte
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return string(commonPrefix)
			}
		}
		commonPrefix = append(commonPrefix, strs[0][i])
	}
	return string(commonPrefix)
}
