package leetcode

import "unicode"

func shortestCompletingWord(licensePlate string, words []string) string {
	lp := genCnter(licensePlate)
	var ret string
	for _, w := range words {
		if match(lp, w) {
			if len(w) < len(ret) || ret == "" {
				ret = w
			}
		}
	}
	return ret
}

func genCnter(lp string) [26]int {
	cnter := [26]int{}
	for _, ch := range lp {
		if unicode.IsLetter(ch) {
			cnter[unicode.ToLower(ch)-'a']++
		}
	}
	return cnter
}

func match(lp [26]int, w string) bool {
	m := [26]int{}
	for _, ch := range w {
		m[ch-'a']++
	}
	for k, v := range lp {
		if m[k] < v {
			return false
		}
	}
	return true
}
