package leetcode

import "strings"

// 解法一 查找时间复杂度 O(1)
type WordFilter struct {
	words map[string]int
}

func Constructor745(words []string) WordFilter {
	wordsMap := make(map[string]int, len(words)*5)
	for k := 0; k < len(words); k++ {
		for i := 0; i <= 10 && i <= len(words[k]); i++ {
			for j := len(words[k]); 0 <= j && len(words[k])-10 <= j; j-- {
				ps := words[k][:i] + "#" + words[k][j:]
				wordsMap[ps] = k
			}
		}
	}
	return WordFilter{words: wordsMap}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	ps := prefix + "#" + suffix
	if index, ok := this.words[ps]; ok {
		return index
	}
	return -1
}

// 解法二 查找时间复杂度 O(N * L)
type WordFilter_ struct {
	input []string
}

func Constructor_745_(words []string) WordFilter_ {
	return WordFilter_{input: words}
}

func (this *WordFilter_) F_(prefix string, suffix string) int {
	for i := len(this.input) - 1; i >= 0; i-- {
		if strings.HasPrefix(this.input[i], prefix) && strings.HasSuffix(this.input[i], suffix) {
			return i
		}
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
