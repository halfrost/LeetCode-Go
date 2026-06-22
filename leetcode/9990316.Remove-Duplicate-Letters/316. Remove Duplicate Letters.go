package leetcode

// 解法一 贪心 + 单调栈。每个字母只保留一个，结果在所有可行解中字典序最小。
// 用 last 记录每个字母最后出现的下标；遍历时若当前字符比栈顶小，且栈顶字符后面还会
// 再出现，就把栈顶弹出，从而让结果尽量小。时间复杂度 O(n)，空间复杂度 O(1)（26 个字母）。
func removeDuplicateLetters(s string) string {
	var last [26]int
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}
	var inStack [26]bool
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if inStack[c-'a'] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > c && last[stack[len(stack)-1]-'a'] > i {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inStack[top-'a'] = false
		}
		stack = append(stack, c)
		inStack[c-'a'] = true
	}
	return string(stack)
}
