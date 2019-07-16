package leetcode

import "strings"

func findOcurrences(text string, first string, second string) []string {
	var res []string
	words := strings.Split(text, " ")
	if len(words) < 3 {
		return []string{}
	}
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			res = append(res, words[i])
		}
	}
	return res
}
