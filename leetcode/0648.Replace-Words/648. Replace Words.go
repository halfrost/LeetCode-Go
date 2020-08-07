package leetcode

import "strings"

// 解法一 哈希表
func replaceWords(dict []string, sentence string) string {
	roots := make(map[byte][]string)
	for _, root := range dict {
		b := root[0]
		roots[b] = append(roots[b], root)
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		b := []byte(word)
		for j := 1; j < len(b) && j <= 100; j++ {
			if findWord(roots, b[0:j]) {
				words[i] = string(b[0:j])
				break
			}
		}
	}
	return strings.Join(words, " ")
}

func findWord(roots map[byte][]string, word []byte) bool {
	if roots[word[0]] == nil {
		return false
	}
	for _, root := range roots[word[0]] {
		if root == string(word) {
			return true
		}
	}
	return false
}

//解法二 Trie
func replaceWords1(dict []string, sentence string) string {
	trie := Constructor208()
	for _, v := range dict {
		trie.Insert(v)
	}
	words := strings.Split(sentence, " ")
	var result []string
	word := ""
	i := 0
	for _, value := range words {
		word = ""
		for i = 1; i < len(value); i++ {
			if trie.Search(value[:i]) {
				word = value[:i]
				break
			}
		}

		if len(word) == 0 {
			result = append(result, value)
		} else {
			result = append(result, word)
		}

	}
	return strings.Join(result, " ")
}

type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor208() Trie {
	return Trie{isWord: false, children: make(map[rune]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{children: make(map[rune]*Trie)}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.isWord
}
