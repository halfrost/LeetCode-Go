# [1653. Minimum Deletions to Make String Balanced](https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/)


## 题目

You are given a string `s` consisting only of characters `'a'` and `'b'`.

You can delete any number of characters in `s` to make `s` **balanced**. `s` is **balanced** if there is no pair of indices `(i,j)` such that `i < j` and `s[i] = 'b'` and `s[j]= 'a'`.

Return *the **minimum** number of deletions needed to make* `s` ***balanced***.

**Example 1:**

```
Input: s = "aababbab"
Output: 2
Explanation: You can either:
Delete the characters at 0-indexed positions 2 and 6 ("aababbab" -> "aaabbb"), or
Delete the characters at 0-indexed positions 3 and 6 ("aababbab" -> "aabbbb").
```

**Example 2:**

```
Input: s = "bbaaaaabb"
Output: 2
Explanation: The only solution is to delete the first two characters.
```

**Constraints:**

- `1 <= s.length <= 105`
- `s[i]` is `'a'` or `'b'`.

## 题目大意

给你一个字符串 s ，它仅包含字符 'a' 和 'b' 。你可以删除 s 中任意数目的字符，使得 s 平衡 。我们称 s 平衡的 当不存在下标对 (i,j) 满足 i < j 且 s[i] = 'b' 同时 s[j]= 'a' 。请你返回使 s 平衡 的 最少 删除次数。

## 解题思路

- 给定一个字符串，要求删除最少次数，使得字母 a 都排在字母 b 的前面。
- 很容易想到的一个解题思路是 DP。定义 `dp[i]` 为字符串下标 [ 0, i ] 这个区间内使得字符串平衡的最少删除次数。当 `s[i] == 'a'` 的时候，有 2 种情况，一种是 `s[i]` 前面全是 `[aa……aa]` 的情况，这个时候只需要把其中的所有的字母 `b` 删除即可。还有一种情况是 `s[i]` 前面有字母 `a` 也有字母 `b`，即 `[aaa……abb……b]`，这种情况就需要考虑 `dp[i-1]` 了。当前字母是 `a`，那么肯定要删除字母 `a`，来维持前面有一段字母 `b` 的情况。当 `s[i] == 'b'` 的时候，不管是 `[aa……aa]` 这种情况，还是 `[aaa……abb……b]` 这种情况，当前字母 `b` 都可以直接附加在后面，也能保证整个字符串是平衡的。所以状态转移方程为 `dp[i+1] = min(dp[i] + 1, bCount), s[i] == 'a'`，`dp[i+1] = dp[i], s[i] == 'b'`。最终答案存在 `dp[n]` 中。由于前后项的递推关系中只用到一次前一项，所以我们还可以优化一下空间，用一个变量保存前一项的结果。优化以后的代码见解法一。
- 这一题还有一个模拟的思路。题目要求找到最小删除字数，那么就是要找到一个“临界点”，在这个临界点的左边删除所有的字母 b，在这个临界点的右边删除所有的字母 a。在所有的“临界点”中找到删除最少的次数。代码实现见解法二。

## 代码

```go
package leetcode

// 解法一 DP
func minimumDeletions(s string) int {
	prev, res, bCount := 0, 0, 0
	for _, c := range s {
		if c == 'a' {
			res = min(prev+1, bCount)
			prev = res
		} else {
			bCount++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 解法二 模拟
func minimumDeletions1(s string) int {
	aCount, bCount, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			aCount++
		}
	}
	res = aCount
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			aCount--
		} else {
			bCount++
		}
		res = min(res, aCount+bCount)
	}
	return res
}
```