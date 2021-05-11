package leetcode

// 解法一 位图
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的bitSet被标记true，说明此字符在X位置重复，需要左侧向前移动，直到将X标记为false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 解法二 滑动窗口-数组桶
func lengthOfLongestSubstring1(s string) int {
	right, left, res := 0, 0, 0
	var indexes [256]int
	for left < len(s) {
		tmp := indexes[s[left]-'a']
		if tmp >= right {
			right = tmp + 1
		}
		indexes[s[left]-'a'] = left
		left++
		res = max(res, left-right)
	}
	return res
}

// 解法二 滑动窗口-哈希桶
func lengthOfLongestSubstring2(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
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
