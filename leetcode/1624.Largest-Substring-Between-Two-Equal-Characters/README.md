# [1624. Largest Substring Between Two Equal Characters](https://leetcode.com/problems/largest-substring-between-two-equal-characters/)

## 题目

Given a string `s`, return *the length of the longest substring between two equal characters, excluding the two characters.* If there is no such substring return `-1`.

A **substring** is a contiguous sequence of characters within a string.

**Example 1:**

```
Input: s = "aa"
Output: 0
Explanation: The optimal substring here is an empty substring between the two 'a's.
```

**Example 2:**

```
Input: s = "abca"
Output: 2
Explanation: The optimal substring here is "bc".
```

**Example 3:**

```
Input: s = "cbzxy"
Output: -1
Explanation: There are no characters that appear twice in s.
```

**Example 4:**

```
Input: s = "cabbac"
Output: 4
Explanation: The optimal substring here is "abba". Other non-optimal substrings include "bb" and "".
```

**Constraints:**

- `1 <= s.length <= 300`
- `s` contains only lowercase English letters.

## 题目大意

给你一个字符串 s，请你返回 两个相同字符之间的最长子字符串的长度 ，计算长度时不含这两个字符。如果不存在这样的子字符串，返回 -1 。子字符串 是字符串中的一个连续字符序列。

## 解题思路

- 简单题。取每个字符，扫描一次字符串，如果在字符串中还能找到相同的字符，则返回最末尾的那个字符，计算这两个字符之间的距离。取最末尾的字符是为了让两个相同的字符距离最长。扫描字符串过程中动态维护最长长度。如果字符串中不存在 2 个相同的字符，则返回 -1 。

## 代码

```go
package leetcode

import "strings"

func maxLengthBetweenEqualCharacters(s string) int {
	res := -1
	for k, v := range s {
		tmp := strings.LastIndex(s, string(v))
		if tmp > 0 {
			if res < tmp-k-1 {
				res = tmp - k - 1
			}
		}
	}
	return res
}
```