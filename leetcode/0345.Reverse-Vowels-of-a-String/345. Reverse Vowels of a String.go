package leetcode

func reverseVowels(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; {
		if !isVowel(b[i]) {
			i++
			continue
		}
		if !isVowel(b[j]) {
			j--
			continue
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func isVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' ||
		s == 'E' || s == 'I' || s == 'O' || s == 'U'
}
