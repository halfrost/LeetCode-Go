package leetcode

func checkInclusion(s1 string, s2 string) bool {
	var freq [256]int
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}
	left, right, count := 0, 0, len(s1)

	for right < len(s2) {
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']--
		right++
		if count == 0 {
			return true
		}
		if right-left == len(s1) {
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}
