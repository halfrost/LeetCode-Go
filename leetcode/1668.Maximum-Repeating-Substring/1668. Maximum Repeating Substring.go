package leetcode

import (
	"strings"
)

func maxRepeating(sequence string, word string) int {
	for i := len(sequence) / len(word); i >= 0; i-- {
		tmp := ""
		for j := 0; j < i; j++ {
			tmp += word
		}
		if strings.Contains(sequence, tmp) {
			return i
		}
	}
	return 0
}
