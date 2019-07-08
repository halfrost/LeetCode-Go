package leetcode

type MagicDictionary struct {
	rdict map[int]string
}

/** Initialize your data structure here. */
func Constructor676() MagicDictionary {
	return MagicDictionary{rdict: make(map[int]string)}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for k, v := range dict {
		this.rdict[k] = v
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	for _, v := range this.rdict {
		n := 0
		if len(word) == len(v) {
			for i := 0; i < len(v); i++ {
				if word[i] != v[i] {
					n += 1
				}
			}
			if n == 1 {
				return true
			}
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
