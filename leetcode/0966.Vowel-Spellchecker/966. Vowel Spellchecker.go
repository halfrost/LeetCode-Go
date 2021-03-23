package leetcode

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	wordsPerfect, wordsCap, wordsVowel := map[string]bool{}, map[string]string{}, map[string]string{}
	for _, word := range wordlist {
		wordsPerfect[word] = true
		wordLow := strings.ToLower(word)
		if _, ok := wordsCap[wordLow]; !ok {
			wordsCap[wordLow] = word
		}
		wordLowVowel := devowel(wordLow)
		if _, ok := wordsVowel[wordLowVowel]; !ok {
			wordsVowel[wordLowVowel] = word
		}
	}
	res, index := make([]string, len(queries)), 0
	for _, query := range queries {
		if _, ok := wordsPerfect[query]; ok {
			res[index] = query
			index++
			continue
		}
		queryL := strings.ToLower(query)
		if v, ok := wordsCap[queryL]; ok {
			res[index] = v
			index++
			continue
		}

		queryLV := devowel(queryL)
		if v, ok := wordsVowel[queryLV]; ok {
			res[index] = v
			index++
			continue
		}
		res[index] = ""
		index++
	}
	return res

}

func devowel(word string) string {
	runes := []rune(word)
	for k, c := range runes {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			runes[k] = '*'
		}
	}
	return string(runes)
}
