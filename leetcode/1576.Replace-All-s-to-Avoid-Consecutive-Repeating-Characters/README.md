# [1576. Replace All ?'s to Avoid Consecutive Repeating Characters](https://leetcode.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/)

## 题目

Given a string `s` containing only lowercase English letters and the `'?'` character, convert **all** the `'?'` characters into lowercase letters such that the final string does not contain any **consecutive repeating** characters. You **cannot** modify the non `'?'` characters.

It is **guaranteed** that there are no consecutive repeating characters in the given string **except** for `'?'`.

Return *the final string after all the conversions (possibly zero) have been made*. If there is more than one solution, return **any of them**. It can be shown that an answer is always possible with the given constraints.

**Example 1:**

```
Input: s = "?zs"
Output: "azs"
Explanation: There are 25 solutions for this problem. From "azs" to "yzs", all are valid. Only "z" is an invalid modification as the string will consist of consecutive repeating characters in "zzs".

```

**Example 2:**

```
Input: s = "ubv?w"
Output: "ubvaw"
Explanation: There are 24 solutions for this problem. Only "v" and "w" are invalid modifications as the strings will consist of consecutive repeating characters in "ubvvw" and "ubvww".

```

**Constraints:**

- `1 <= s.length <= 100`
- `s` consist of lowercase English letters and `'?'`.

## 题目大意

给你一个仅包含小写英文字母和 '?' 字符的字符串 s，请你将所有的 '?' 转换为若干小写字母，使最终的字符串不包含任何 连续重复 的字符。注意：你 不能 修改非 '?' 字符。

题目测试用例保证 除 '?' 字符 之外，不存在连续重复的字符。在完成所有转换（可能无需转换）后返回最终的字符串。如果有多个解决方案，请返回其中任何一个。可以证明，在给定的约束条件下，答案总是存在的。

## 解题思路

- 简单题。找到源字符串中 ‘?’ 字符的位置，然后依次用 a ~ z 字符去替换，替换进去的字符不能和前后字符相同即可。

## 代码

```go
package leetcode

func modifyString(s string) string {
	res := []byte(s)
	for i, ch := range res {
		if ch == '?' {
			for b := byte('a'); b <= 'z'; b++ {
				if !(i > 0 && res[i-1] == b || i < len(res)-1 && res[i+1] == b) {
					res[i] = b
					break
				}
			}
		}
	}
	return string(res)
}
```