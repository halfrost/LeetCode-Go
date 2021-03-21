package leetcode

import "fmt"

func reorderedPowerOf2(n int) bool {
	sample, i := fmt.Sprintf("%v", n), 1
	for len(fmt.Sprintf("%v", i)) <= len(sample) {
		t := fmt.Sprintf("%v", i)
		if len(t) == len(sample) && isSame(t, sample) {
			return true
		}
		i = i << 1
	}
	return false
}

func isSame(t, s string) bool {
	m := make(map[rune]int)
	for _, v := range t {
		m[v]++
	}
	for _, v := range s {
		m[v]--
		if m[v] < 0 {
			return false
		}
		if m[v] == 0 {
			delete(m, v)
		}
	}
	return len(m) == 0
}
