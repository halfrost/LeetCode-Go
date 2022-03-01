# [2182. Construct String With Repeat Limit](https://leetcode.com/problems/construct-string-with-repeat-limit/)


## 题目

You are given a string `s` and an integer `repeatLimit`. Construct a new string `repeatLimitedString` using the characters of `s` such that no letter appears **more than** `repeatLimit` times **in a row**. You do **not** have to use all characters from `s`.

Return *the **lexicographically largest*** `repeatLimitedString` *possible*.

A string `a` is **lexicographically larger** than a string `b` if in the first position where `a` and `b` differ, string `a` has a letter that appears later in the alphabet than the corresponding letter in `b`. If the first `min(a.length, b.length)` characters do not differ, then the longer string is the lexicographically larger one.

**Example 1:**

```
Input: s = "cczazcc", repeatLimit = 3
Output: "zzcccac"
Explanation: We use all of the characters from s to construct the repeatLimitedString "zzcccac".
The letter 'a' appears at most 1 time in a row.
The letter 'c' appears at most 3 times in a row.
The letter 'z' appears at most 2 times in a row.
Hence, no letter appears more than repeatLimit times in a row and the string is a valid repeatLimitedString.
The string is the lexicographically largest repeatLimitedString possible so we return "zzcccac".
Note that the string "zzcccca" is lexicographically larger but the letter 'c' appears more than 3 times in a row, so it is not a valid repeatLimitedString.

```

**Example 2:**

```
Input: s = "aababab", repeatLimit = 2
Output: "bbabaa"
Explanation: We use only some of the characters from s to construct the repeatLimitedString "bbabaa".
The letter 'a' appears at most 2 times in a row.
The letter 'b' appears at most 2 times in a row.
Hence, no letter appears more than repeatLimit times in a row and the string is a valid repeatLimitedString.
The string is the lexicographically largest repeatLimitedString possible so we return "bbabaa".
Note that the string "bbabaaa" is lexicographically larger but the letter 'a' appears more than 2 times in a row, so it is not a valid repeatLimitedString.

```

**Constraints:**

- `1 <= repeatLimit <= s.length <= 10^5`
- `s` consists of lowercase English letters.

## 题目大意

给你一个字符串 s 和一个整数 repeatLimit ，用 s 中的字符构造一个新字符串 repeatLimitedString ，使任何字母 连续 出现的次数都不超过 repeatLimit 次。你不必使用 s 中的全部字符。

返回 字典序最大的 repeatLimitedString 。

如果在字符串 a 和 b 不同的第一个位置，字符串 a 中的字母在字母表中出现时间比字符串 b 对应的字母晚，则认为字符串 a 比字符串 b 字典序更大 。如果字符串中前 min(a.length, b.length) 个字符都相同，那么较长的字符串字典序更大。

## 解题思路

- 利用贪心的思想，由于题意要求返回字典序最大的字符串，所以先从字典序最大的字母开始选起。然后选择当前字典序最大的字母个数和 limit 的最小值。如果当前字典序最大的字母比较多，多于 limit，不能一直选择它。选完 limit 个以后，需要选一个字典序次大的字母，选完这个字母以后再次选择字典序最大的字母。因为 limit 限制字母不能连续多于 limit 个。如此循环，直到所有的字母都选完。这样的策略排列出来的字母串为最大字典序。

## 代码

```go
package leetcode

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[int(c-'a')]++
	}
	var ns []byte
	for i := 25; i >= 0; {
		k := i - 1
		for cnt[i] > 0 {
			for j := 0; j < min(cnt[i], repeatLimit); j++ {
				ns = append(ns, byte(i)+'a')
			}
			cnt[i] -= repeatLimit
			if cnt[i] > 0 {
				for ; k >= 0 && cnt[k] == 0; k-- {
				}
				if k < 0 {
					break
				} else {
					ns = append(ns, byte(k)+'a')
					cnt[k]--
				}
			}
		}
		i = k
	}
	return string(ns)
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
```