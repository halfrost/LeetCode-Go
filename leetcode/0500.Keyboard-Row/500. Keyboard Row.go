package leetcode

import "strings"

func findWords500(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	output := make([]string, 0)
	for _, s := range words {
		if len(s) == 0 {
			continue
		}
		lowerS := strings.ToLower(s)
		oneRow := false
		for _, r := range rows {
			if strings.ContainsAny(lowerS, r) {
				oneRow = !oneRow
				if !oneRow {
					break
				}
			}
		}
		if oneRow {
			output = append(output, s)
		}
	}
	return output
}
