# [1143. Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/)

## 题目

Given two strings `text1` and `text2`, return *the length of their longest **common subsequence**.* If there is no **common subsequence**, return `0`.

A **subsequence** of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

- For example, `"ace"` is a subsequence of `"abcde"`.

A **common subsequence** of two strings is a subsequence that is common to both strings.

**Example 1:**

```
Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
```

**Example 2:**

```
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
```

**Example 3:**

```
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
```

**Constraints:**

- `1 <= text1.length, text2.length <= 1000`
- `text1` and `text2` consist of only lowercase English characters.

## 题目大意

给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

## 解题思路

- 这一题是经典的最长公共子序列的问题。解题思路是二维动态规划。假设字符串 `text1` 和 `text2` 的长度分别为 `m` 和 `n`，创建 `m+1` 行 `n+1` 列的二维数组 `dp`，定义 `dp[i][j]` 表示长度为 i 的 `text1[0:i-1]` 和长度为 j 的 `text2[0:j-1]` 的最长公共子序列的长度。先考虑边界条件。当 `i = 0` 时，`text1[]` 为空字符串，它与任何字符串的最长公共子序列的长度都是 `0`，所以 `dp[0][j] = 0`。同理当 `j = 0` 时，`text2[]` 为空字符串，它与任何字符串的最长公共子序列的长度都是 `0`，所以 `dp[i][0] = 0`。由于二维数组的大小特意增加了 `1`，即 `m+1` 和 `n+1`，并且默认值是 `0`，所以不需要再初始化赋值了。
- 当 `text1[i−1] = text2[j−1]` 时，将这两个相同的字符称为公共字符，考虑 `text1[0:i−1]` 和 `text2[0:j−1]` 的最长公共子序列，再增加一个字符（即公共字符）即可得到 `text1[0:i]` 和 `text2[0:j]` 的最长公共子序列，所以 `dp[i][j]=dp[i−1][j−1]+1`。当 `text1[i−1] != text2[j−1]` 时，最长公共子序列一定在 `text[0:i-1], text2[0:j]` 和 `text[0:i], text2[0:j-1]` 中取得。即 `dp[i][j] = max(dp[i-1][j], dp[i][j-1])`。所以状态转移方程如下：

    $$dp[i][j] = \left\{\begin{matrix}dp[i-1][j-1]+1 &,text1[i-1]=text2[j-1]\\max(dp[i-1][j],dp[i][j-1])&,text1[i-1]\neq text2[j-1]\end{matrix}\right.$$

- 最终结果存储在 `dp[len(text1)][len(text2)]` 中。时间复杂度 `O(mn)`，空间复杂度 `O(mn)`，其中 `m` 和 `n` 分别是 `text1` 和 `text2` 的长度。

## 代码

```go
package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```