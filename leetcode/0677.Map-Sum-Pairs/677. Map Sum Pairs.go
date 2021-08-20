package leetcode

type MapSum struct {
	keys map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.keys[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	prefixAsRunes, res := []rune(prefix), 0
	for key, val := range this.keys {
		if len(key) >= len(prefix) {
			shouldSum := true
			for i, char := range key {
				if i >= len(prefixAsRunes) {
					break
				}
				if prefixAsRunes[i] != char {
					shouldSum = false
					break
				}
			}
			if shouldSum {
				res += val
			}
		}
	}
	return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
