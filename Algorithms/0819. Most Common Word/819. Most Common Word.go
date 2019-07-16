package leetcode

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	freqMap, start := make(map[string]int), -1
	for i, c := range paragraph {
		if c == ' ' || c == '!' || c == '?' || c == '\'' || c == ',' || c == ';' || c == '.' {
			if start > -1 {
				word := strings.ToLower(paragraph[start:i])
				freqMap[word]++
			}
			start = -1
		} else {
			if start == -1 {
				start = i
			}
		}
	}
	if start != -1 {
		word := strings.ToLower(paragraph[start:])
		freqMap[word]++
	}
	// Strip the banned words from the freqmap
	for _, bannedWord := range banned {
		delete(freqMap, bannedWord)
	}
	// Find most freq word
	mostFreqWord, mostFreqCount := "", 0
	for word, freq := range freqMap {
		if freq > mostFreqCount {
			mostFreqWord = word
			mostFreqCount = freq
		}
	}
	return mostFreqWord
}
