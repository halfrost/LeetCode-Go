package leetcode

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) || len(s) <= 1 {
		return false
	}
	mp := make(map[byte]int)
	if s == goal {
		for i := 0; i < len(s); i++ {
			if _, ok := mp[s[i]]; ok {
				return true
			}
			mp[s[i]]++
		}
		return false
	}
	first, second := -1, -1
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}
