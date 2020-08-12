# [67. Add Binary](https://leetcode.com/problems/add-binary/)


## 题目

Given two binary strings, return their sum (also a binary string).

The input strings are both **non-empty** and contains only characters `1` or `0`.

**Example 1**:

```
Input: a = "11", b = "1"
Output: "100"
```

**Example 2**:

```
Input: a = "1010", b = "1011"
Output: "10101"
```

## 题目大意

给你两个二进制字符串，返回它们的和（用二进制表示）。输入为 非空 字符串且只包含数字 1 和 0。

## 解题思路

- 要求输出 2 个二进制数的和，结果也用二进制表示。
- 简单题。按照二进制的加法规则做加法即可。

## 代码

```go

package leetcode

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	res := make([]string, len(a)+1)
	i, j, k, c := len(a)-1, len(b)-1, len(a), 0
	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))
		res[k] = strconv.Itoa((ai + bj + c) % 2)
		c = (ai + bj + c) / 2
		i--
		j--
		k--
	}

	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		res[k] = strconv.Itoa((ai + c) % 2)
		c = (ai + c) / 2
		i--
		k--
	}

	if c > 0 {
		res[k] = strconv.Itoa(c)
	}

	return strings.Join(res, "")
}

```