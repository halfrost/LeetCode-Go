# [1614. Maximum Nesting Depth of the Parentheses](https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/)

## 题目

A string is a **valid parentheses string** (denoted **VPS**) if it meets one of the following:

- It is an empty string `""`, or a single character not equal to `"("` or `")"`,
- It can be written as `AB` (`A` concatenated with `B`), where `A` and `B` are **VPS**'s, or
- It can be written as `(A)`, where `A` is a **VPS**.

We can similarly define the **nesting depth** `depth(S)` of any VPS `S` as follows:

- `depth("") = 0`
- `depth(C) = 0`, where `C` is a string with a single character not equal to `"("` or `")"`.
- `depth(A + B) = max(depth(A), depth(B))`, where `A` and `B` are **VPS**'s.
- `depth("(" + A + ")") = 1 + depth(A)`, where `A` is a **VPS**.

For example, `""`, `"()()"`, and `"()(()())"` are **VPS**'s (with nesting depths 0, 1, and 2), and `")("` and `"(()"` are not **VPS**'s.

Given a **VPS** represented as string `s`, return *the **nesting depth** of* `s`.

**Example 1:**

```
Input: s = "(1+(2*3)+((8)/4))+1"
Output: 3
Explanation: Digit 8 is inside of 3 nested parentheses in the string.
```

**Example 2:**

```
Input: s = "(1)+((2))+(((3)))"
Output: 3
```

**Example 3:**

```
Input: s = "1+(2*3)/(2-1)"
Output: 1
```

**Example 4:**

```
Input: s = "1"
Output: 0
```

**Constraints:**

- `1 <= s.length <= 100`
- `s` consists of digits `0-9` and characters `'+'`, `'-'`, `'*'`, `'/'`, `'('`, and `')'`.
- It is guaranteed that parentheses expression `s` is a **VPS**.

## 题目大意

如果字符串满足以下条件之一，则可以称之为 有效括号字符串（valid parentheses string，可以简写为 VPS）：

- 字符串是一个空字符串 ""，或者是一个不为 "(" 或 ")" 的单字符。
- 字符串可以写为 AB（A 与 B 字符串连接），其中 A 和 B 都是 有效括号字符串 。
- 字符串可以写为 (A)，其中 A 是一个 有效括号字符串 。

类似地，可以定义任何有效括号字符串 S 的 嵌套深度 depth(S)：

- depth("") = 0
- depth(C) = 0，其中 C 是单个字符的字符串，且该字符不是 "(" 或者 ")"
- depth(A + B) = max(depth(A), depth(B))，其中 A 和 B 都是 有效括号字符串
- depth("(" + A + ")") = 1 + depth(A)，其中 A 是一个 有效括号字符串

例如：""、"()()"、"()(()())" 都是 有效括号字符串（嵌套深度分别为 0、1、2），而 ")(" 、"(()" 都不是 有效括号字符串 。给你一个 有效括号字符串 s，返回该字符串的 s 嵌套深度 。

## 解题思路

- 简单题。求一个括号字符串的嵌套深度。题目给的括号字符串都是有效的，所以不需要考虑非法的情况。扫描一遍括号字符串，遇到 `(` 可以无脑 ++，并动态维护最大值，遇到 `)` 可以无脑 - - 。最后输出最大值即可。

## 代码

```go
package leetcode

func maxDepth(s string) int {
	res, cur := 0, 0
	for _, c := range s {
		if c == '(' {
			cur++
			res = max(res, cur)
		} else if c == ')' {
			cur--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```