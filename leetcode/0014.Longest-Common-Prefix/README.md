# [14. Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/)

## 题目

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

**Example 1**:

    Input: strs = ["flower","flow","flight"]
    Output: "fl"

**Example 2**:

    Input: strs = ["dog","racecar","car"]
    Output: ""
    Explanation: There is no common prefix among the input strings.

**Constraints:**

- 1 <= strs.length <= 200
- 0 <= strs[i].length <= 200
- strs[i] consists of only lower-case English letters.

## 题目大意

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

## 解题思路

- 对 strs 按照字符串长度进行升序排序，求出 strs 中长度最小字符串的长度 minLen
- 逐个比较长度最小字符串与其它字符串中的字符，如果不相等就返回 commonPrefix,否则就把该字符加入 commonPrefix

## 代码

```go

package leetcode

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}

	return prefix
}
```
