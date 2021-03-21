# [1663. Smallest String With A Given Numeric Value](https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/)

## 题目

The **numeric value** of a **lowercase character** is defined as its position `(1-indexed)` in the alphabet, so the numeric value of `a` is `1`, the numeric value of `b` is `2`, the numeric value of `c` is `3`, and so on.

The **numeric value** of a **string** consisting of lowercase characters is defined as the sum of its characters' numeric values. For example, the numeric value of the string `"abe"` is equal to `1 + 2 + 5 = 8`.

You are given two integers `n` and `k`. Return *the **lexicographically smallest string** with **length** equal to `n` and **numeric value** equal to `k`.*

Note that a string `x` is lexicographically smaller than string `y` if `x` comes before `y` in dictionary order, that is, either `x` is a prefix of `y`, or if `i` is the first position such that `x[i] != y[i]`, then `x[i]` comes before `y[i]` in alphabetic order.

**Example 1:**

```
Input: n = 3, k = 27
Output: "aay"
Explanation: The numeric value of the string is 1 + 1 + 25 = 27, and it is the smallest string with such a value and length equal to 3.
```

**Example 2:**

```
Input: n = 5, k = 73
Output: "aaszz"
```

**Constraints:**

- `1 <= n <= 105`
- `n <= k <= 26 * n`

## 题目大意

小写字符 的 数值 是它在字母表中的位置（从 1 开始），因此 a 的数值为 1 ，b 的数值为 2 ，c 的数值为 3 ，以此类推。字符串由若干小写字符组成，字符串的数值 为各字符的数值之和。例如，字符串 "abe" 的数值等于 1 + 2 + 5 = 8 。给你两个整数 n 和 k 。返回 长度 等于 n 且 数值 等于 k 的 字典序最小 的字符串。注意，如果字符串 x 在字典排序中位于 y 之前，就认为 x 字典序比 y 小，有以下两种情况：

- x 是 y 的一个前缀；
- 如果 i 是 x[i] != y[i] 的第一个位置，且 x[i] 在字母表中的位置比 y[i] 靠前。

## 解题思路

- 给出 n 和 k，要求找到字符串长度为 n，字母在字母表内位置总和为 k 的最小字典序字符串。
- 这一题笔者读完题，比赛的时候直接用 DFS 撸了一版。赛后看了时间复杂度马马虎虎，感觉还有优化的空间。DFS 会遍历出所有的解，实际上这一题只要求最小字典序，所以 DFS 剪枝的时候要加上判断字典序的判断，如果新添加进来的字母比已经保存的字符串的相应位置上的字母字典序大，那么就直接 return，这个答案一定不会是最小字典序。代码见解法二
- 想到这里，其实 DFS 不必要，直接用 for 循环就可找到最小字典序的字符串。代码见解法一。

## 代码

```go
package leetcode

// 解法一 贪心
func getSmallestString(n int, k int) string {
    str, i, j := make([]byte, n), 0, 0
    for i = n-1; i <= k-26; i, k = i-1, k-26 {
        str[i] = 'z'
    }
    if i >= 0 {
        str[i] = byte('a' + k-1-i)
        for ; j < i; j++ {
            str[j] = 'a'
        }
    }
    return string(str)
}

// 解法二 DFS
func getSmallestString1(n int, k int) string {
	if n == 0 {
		return ""
	}
	res, c := "", []byte{}
	findSmallestString(0, n, k, 0, c, &res)
	return res
}

func findSmallestString(value int, length, k, index int, str []byte, res *string) {
	if len(str) == length && value == k {
		tmp := string(str)
		if (*res) == "" {
			*res = tmp
		}
		if tmp < *res && *res != "" {
			*res = tmp
		}
		return
	}
	if len(str) >= index && (*res) != "" && str[index-1] > (*res)[index-1] {
		return
	}
	for j := 0; j < 26; j++ {
		if k-value > (length-len(str))*26 || value > k {
			return
		}
		str = append(str, byte(int('a')+j))
		value += j + 1
		findSmallestString(value, length, k, index+1, str, res)
		str = str[:len(str)-1]
		value -= j + 1

	}
}
```