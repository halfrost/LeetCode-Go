package leetcode

import "strings"

func findDuplicate(paths []string) [][]string {
	cache := make(map[string][]string)
	for _, path := range paths {
		parts := strings.Split(path, " ")
		dir := parts[0]
		for i := 1; i < len(parts); i++ {
			bracketPosition := strings.IndexByte(parts[i], '(')
			content := parts[i][bracketPosition+1 : len(parts[i])-1]
			cache[content] = append(cache[content], dir+"/"+parts[i][:bracketPosition])
		}
	}
	res := make([][]string, 0, len(cache))
	for _, group := range cache {
		if len(group) >= 2 {
			res = append(res, group)
		}
	}
	return res
}
