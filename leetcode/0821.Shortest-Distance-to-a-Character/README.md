# [821. Shortest Distance to a Character](https://leetcode.com/problems/shortest-distance-to-a-character/)

## 题目

Given a string `s` and a character `c` that occurs in `s`, return *an array of integers `answer` where* `answer.length == s.length` *and* `answer[i]` *is the shortest distance from* `s[i]` *to the character* `c` *in* `s`.

**Example 1:**

```
Input: s = "loveleetcode", c = "e"
Output: [3,2,1,0,1,0,0,1,2,2,1,0]
```

**Example 2:**

```
Input: s = "aaab", c = "b"
Output: [3,2,1,0]
```

**Constraints:**

- `1 <= s.length <= 104`
- `s[i]` and `c` are lowercase English letters.
- `c` occurs at least once in `s`.

## 题目大意

给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。

## 解题思路

- 解法一：从左至右更新一遍到 C 的值距离，再从右至左更新一遍到 C 的值，取两者中的最小值。
- 解法二：依次扫描字符串 S，针对每一个非字符 C 的字符，分别往左扫一次，往右扫一次，计算出距离目标字符 C 的距离，然后取左右两个距离的最小值存入最终答案数组中。

## 代码

```go
package leetcode

import (
	"math"
)

// 解法一
func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)
	for i := range res {
		res[i] = n
	}
	for i := 0; i < n; i++ {
		if s[i] == c {
			res[i] = 0
		} else if i > 0 {
			res[i] = res[i-1] + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && res[i+1]+1 < res[i] {
			res[i] = res[i+1] + 1
		}
	}
	return res
}

// 解法二
func shortestToChar1(s string, c byte) []int {
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			res[i] = 0
		} else {
			left, right := math.MaxInt32, math.MaxInt32
			for j := i + 1; j < len(s); j++ {
				if s[j] == c {
					right = j - i
					break
				}
			}
			for k := i - 1; k >= 0; k-- {
				if s[k] == c {
					left = i - k
					break
				}
			}
			res[i] = min(left, right)
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```