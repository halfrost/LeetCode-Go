package leetcode

import (
	"strings"
)

func maxRepeating(sequence string, word string) int {
	res := 0
	for i := len(sequence) / len(word); i >= 1; i-- {
		tmp := ""
		for j := 0; j < i; j++ {
			tmp += word
		}
		if strings.Contains(sequence, tmp) {
			res = i
			break
		}
	}
	return res
}
