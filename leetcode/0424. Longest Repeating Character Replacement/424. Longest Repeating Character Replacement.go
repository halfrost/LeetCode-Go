package leetcode

func characterReplacement(s string, k int) int {
	res, left, right, counter, freq := 0, 0, 1, 0, make([]int, 26)
	for right = 0; right < len(s); right++ {
		freq[s[right]-'A']++
		counter = max(counter, freq[s[right]-'A'])
		for right-left+1-counter > k {
			freq[s[left]-'A']--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}
