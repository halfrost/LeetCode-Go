# [696. Count Binary Substrings](https://leetcode.com/problems/count-binary-substrings/)


## 题目

Give a string `s`, count the number of non-empty (contiguous) substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.

Substrings that occur multiple times are counted the number of times they occur.

**Example 1:**

```
Input: "00110011"
Output: 6
Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".

Notice that some of these substrings repeat and are counted the number of times they occur.

Also, "00110011" is not a valid substring becauseall the 0's (and 1's) are not grouped together.

```

**Example 2:**

```
Input: "10101"
Output: 4
Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.

```

**Note:**

- `s.length` will be between 1 and 50,000.
- `s` will only consist of "0" or "1" characters.

## 题目大意

给定一个字符串 s，计算具有相同数量 0 和 1 的非空（连续）子字符串的数量，并且这些子字符串中的所有 0 和所有 1 都是连续的。重复出现的子串要计算它们出现的次数。

## 解题思路

- 简单题。先分组统计 0 和 1 的个数，例如，`0110001111` 按照 0 和 1 分组统计出来的结果是 [1, 2, 3, 4]。再拼凑结果。相邻 2 组取两者最短的，例如 `0110001111`，凑成的结果应该是 min(1,2)，min(2,3)，min(3,4)，即 `01`，`01`，`10`，`1100`，`0011`，`000111`。时间复杂度 O(n)，空间复杂度 O(1)。

## 代码

```go
package leetcode

func countBinarySubstrings(s string) int {
	last, res := 0, 0
	for i := 0; i < len(s); {
		c, count := s[i], 1
		for i++; i < len(s) && s[i] == c; i++ {
			count++
		}
		res += min(count, last)
		last = count
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```