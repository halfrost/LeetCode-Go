package leetcode

// 解法一 滑动窗口-数组桶
func lengthOfLongestSubstring1(s string) int {
	right, left, res := 0, 0, 0
	var m [256]int
	for left < len(s) {
		tmp := m[s[left]-'a']
		if tmp >= right {
			right = tmp + 1
		}
		m[s[left]-'a'] = left
		left++
		res = max(res, left-right)
	}
	return res
}

// 解法二 滑动窗口-哈希桶
func lengthOfLongestSubstring(s string) int {
	right, left, res := 0, 0, 0
	m := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := m[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		m[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
