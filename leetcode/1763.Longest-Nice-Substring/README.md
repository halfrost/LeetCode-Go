# [1763. Longest Nice Substring](https://leetcode.com/problems/longest-nice-substring/)


## 题目

A string `s` is **nice** if, for every letter of the alphabet that `s` contains, it appears **both** in uppercase and lowercase. For example, `"abABB"` is nice because `'A'` and `'a'` appear, and `'B'` and `'b'` appear. However, `"abA"` is not because `'b'` appears, but `'B'` does not.

Given a string `s`, return *the longest **substring** of `s` that is **nice**. If there are multiple, return the substring of the **earliest** occurrence. If there are none, return an empty string*.

**Example 1:**

```
Input: s = "YazaAay"
Output: "aAa"
Explanation:"aAa" is a nice string because 'A/a' is the only letter of the alphabet in s, and both 'A' and 'a' appear.
"aAa" is the longest nice substring.

```

**Example 2:**

```
Input: s = "Bb"
Output: "Bb"
Explanation: "Bb" is a nice string because both 'B' and 'b' appear. The whole string is a substring.

```

**Example 3:**

```
Input: s = "c"
Output: ""
Explanation: There are no nice substrings.

```

**Constraints:**

- `1 <= s.length <= 100`
- `s` consists of uppercase and lowercase English letters.

## 题目大意

当一个字符串 s 包含的每一种字母的大写和小写形式 同时 出现在 s 中，就称这个字符串 s 是 美好 字符串。比方说，"abABB" 是美好字符串，因为 'A' 和 'a' 同时出现了，且 'B' 和 'b' 也同时出现了。然而，"abA" 不是美好字符串因为 'b' 出现了，而 'B' 没有出现。

给你一个字符串 s ，请你返回 s 最长的 美好子字符串 。如果有多个答案，请你返回 最早 出现的一个。如果不存在美好子字符串，请你返回一个空字符串。

## 解题思路

- 解法一，暴力解法。枚举每一段字符串，判断这个子字符串内是否满足美好字符串的定义，即字母的大小写是否同时出现。
- 解法二，这个解法是解法一的小幅优化版，利用二进制记录状态。先构造二进制状态串，再利用直接比较这个二进制串。
- 解法三，分治。以 `i` 为分割点依次切开字符串。左右两个字符串分别判断是否满足美好字符串的定义。左右分开的字符串还可以继续划分。直至分到一个字母为止。在这个过程中记录最早出现的字符串。

## 代码

```go
package leetcode

import "unicode"

// 解法一 分治，时间复杂度 O(n)
func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}

	chars := map[rune]int{}
	for _, r := range s {
		chars[r]++
	}

	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		_, u := chars[unicode.ToUpper(r)]
		_, l := chars[unicode.ToLower(r)]
		if u && l {
			continue
		}
		left := longestNiceSubstring(s[:i])
		right := longestNiceSubstring(s[i+1:])
		if len(left) >= len(right) {
			return left
		} else {
			return right
		}
	}
	return s
}

// 解法二 用二进制表示状态
func longestNiceSubstring1(s string) (ans string) {
	for i := range s {
		lower, upper := 0, 0
		for j := i; j < len(s); j++ {
			if unicode.IsLower(rune(s[j])) {
				lower |= 1 << (s[j] - 'a')
			} else {
				upper |= 1 << (s[j] - 'A')
			}
			if lower == upper && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return
}

// 解法三 暴力枚举，时间复杂度 O(n^2)
func longestNiceSubstring2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		m := map[byte]int{}
		m[s[i]]++
		for j := i + 1; j < len(s); j++ {
			m[s[j]]++
			if checkNiceString(m) && (j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func checkNiceString(m map[byte]int) bool {
	for k := range m {
		if k >= 97 && k <= 122 {
			if _, ok := m[k-32]; !ok {
				return false
			}
		}
		if k >= 65 && k <= 90 {
			if _, ok := m[k+32]; !ok {
				return false
			}
		}
	}
	return true
}
```