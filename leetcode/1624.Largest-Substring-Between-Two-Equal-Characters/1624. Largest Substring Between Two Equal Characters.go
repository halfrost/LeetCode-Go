package leetcode

import "strings"

func maxLengthBetweenEqualCharacters(s string) int {
	res := -1
	for k, v := range s {
		tmp := strings.LastIndex(s, string(v))
		if tmp > 0 {
			if res < tmp-k-1 {
				res = tmp - k - 1
			}
		}
	}
	return res
}
