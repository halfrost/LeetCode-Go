# [58. Length of Last Word](https://leetcode.com/problems/length-of-last-word/)


## 题目

Given a string `s` consisting of some words separated by some number of spaces, return *the length of the **last** word in the string.*

A **word** is a maximal substring consisting of non-space characters only.

**Example 1:**

```
Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.

```

**Example 2:**

```
Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.

```

**Example 3:**

```
Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.

```

**Constraints:**

- `1 <= s.length <= 104`
- `s` consists of only English letters and spaces `' '`.
- There will be at least one word in `s`.

## 题目大意

给你一个字符串 `s`，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。**单词** 是指仅由字母组成、不包含任何空格字符的最大子字符串。

## 解题思路

- 先从后过滤掉空格找到单词尾部，再从尾部向前遍历，找到单词头部，最后两者相减，即为单词的长度。

## 代码

```go
package leetcode

func lengthOfLastWord(s string) int {
	last := len(s) - 1
	for last >= 0 && s[last] == ' ' {
		last--
	}
	if last < 0 {
		return 0
	}
	first := last
	for first >= 0 && s[first] != ' ' {
		first--
	}
	return last - first
}
```