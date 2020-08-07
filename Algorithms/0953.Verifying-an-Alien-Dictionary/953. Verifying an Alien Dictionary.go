package leetcode

func isAlienSorted(words []string, order string) bool {
	if len(words) < 2 {
		return true
	}
	hash := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		hash[order[i]] = i
	}
	for i := 0; i < len(words)-1; i++ {
		pointer, word, wordplus := 0, words[i], words[i+1]
		for pointer < len(word) && pointer < len(wordplus) {
			if hash[word[pointer]] > hash[wordplus[pointer]] {
				return false
			}
			if hash[word[pointer]] < hash[wordplus[pointer]] {
				break
			} else {
				pointer = pointer + 1
			}
		}
		if pointer < len(word) && pointer >= len(wordplus) {
			return false
		}
	}
	return true
}
