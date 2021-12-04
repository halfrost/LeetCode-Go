package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	var cnt [26]int
	for _, v := range magazine {
		cnt[v-'a']++
	}
	for _, v := range ransomNote {
		cnt[v-'a']--
		if cnt[v-'a'] < 0 {
			return false
		}
	}
	return true
}
