# [520. Detect Capital](https://leetcode.com/problems/detect-capital/)


## 题目

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".

All letters in this word are not capitals, like "leetcode".

Only the first letter in this word is capital, like "Google".

Given a string word, return true if the usage of capitals in it is right.

**Example 1:**

```
Input: word = "USA"
Output: true
```

**Example 2:**

```
Input: word = "FlaG"
Output: false
```

**Constraints:**

- 1 <= word.length <= 100
- word consists of lowercase and uppercase English letters.

## 题目大意

我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如 "USA" 。
单词中所有字母都不是大写，比如 "leetcode" 。
如果单词不只含有一个字母，只有首字母大写，比如"Google" 。

给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。

## 解题思路

- 把 word 分别转换为全部小写 wLower，全部大写 wUpper，首字母大写的字符串 wCaptial
- 判断 word 是否等于 wLower, wUpper, wCaptial 中的一个，如果是返回 true，否则返回 false

## 代码

```go

package leetcode

import "strings"

func detectCapitalUse(word string) bool {
	wLower := strings.ToLower(word)
	wUpper := strings.ToUpper(word)
	wCaptial := strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
	if wCaptial == word || wLower == word || wUpper == word {
		return true
	}
	return false
}

```