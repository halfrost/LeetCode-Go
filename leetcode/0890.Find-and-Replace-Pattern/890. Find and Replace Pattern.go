package leetcode

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for _, word := range words {
		if match(word, pattern) {
			res = append(res, word)
		}
	}
	return res
}

func match(w, p string) bool {
	if len(w) != len(p) {
		return false
	}
	m, used := make(map[uint8]uint8), make(map[uint8]bool)
	for i := 0; i < len(w); i++ {
		if v, ok := m[p[i]]; ok {
			if w[i] != v {
				return false
			}
		} else {
			if used[w[i]] {
				return false
			}
			m[p[i]] = w[i]
			used[w[i]] = true
		}
	}
	return true
}
