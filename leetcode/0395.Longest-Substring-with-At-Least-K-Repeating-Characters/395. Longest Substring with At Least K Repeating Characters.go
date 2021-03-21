package leetcode

import "strings"

// 解法一 滑动窗口
func longestSubstring(s string, k int) int {
	res := 0
	for t := 1; t <= 26; t++ {
		freq, total, lessK, left, right := [26]int{}, 0, 0, 0, -1
		for left < len(s) {
			if right+1 < len(s) && total <= t {
				if freq[s[right+1]-'a'] == 0 {
					total++
					lessK++
				}
				freq[s[right+1]-'a']++
				if freq[s[right+1]-'a'] == k {
					lessK--
				}
				right++
			} else {
				if freq[s[left]-'a'] == k {
					lessK++
				}
				freq[s[left]-'a']--
				if freq[s[left]-'a'] == 0 {
					total--
					lessK--
				}
				left++
			}
			if lessK == 0 {
				res = max(res, right-left+1)
			}

		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法二 递归分治
func longestSubstring1(s string, k int) int {
	if s == "" {
		return 0
	}
	freq, split, res := [26]int{}, byte(0), 0
	for _, ch := range s {
		freq[ch-'a']++
	}
	for i, c := range freq[:] {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}
	for _, subStr := range strings.Split(s, string(split)) {
		res = max(res, longestSubstring1(subStr, k))
	}
	return res
}
