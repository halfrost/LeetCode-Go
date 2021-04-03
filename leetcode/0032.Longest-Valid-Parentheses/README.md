# [32. Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/)


## 题目

Given a string containing just the characters `'('` and `')'`, find the length of the longest valid (well-formed) parentheses substring.

**Example 1:**

```
Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
```

**Example 2:**

```
Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
```

**Example 3:**

```
Input: s = ""
Output: 0
```

**Constraints:**

- `0 <= s.length <= 3 * 104`
- `s[i]` is `'('`, or `')'`.

## 题目大意

给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

## 解题思路

- 提到括号匹配，第一时间能让人想到的就是利用栈。这里需要计算嵌套括号的总长度，所以栈里面不能单纯的存左括号，而应该存左括号在原字符串的下标，这样通过下标相减可以获取长度。那么栈如果是非空，栈底永远存的是当前遍历过的字符串中**上一个没有被匹配的右括号的下标**。**上一个没有被匹配的右括号的下标**可以理解为每段括号匹配之间的“隔板”。例如，`())((()))`，第三个右括号，即为左右 2 段正确的括号匹配中间的“隔板”。“隔板”的存在影响计算最长括号长度。如果不存在“隔板”，前后 2 段正确的括号匹配应该“融合”在一起，最长长度为 `2 + 6 = 8`，但是这里存在了“隔板”，所以最长长度仅为 `6`。
- 具体算法实现，遇到每个 `'('` ，将它的下标压入栈中。对于遇到的每个 `')'`，先弹出栈顶元素表示匹配了当前右括号。如果栈为空，说明当前的右括号为没有被匹配的右括号，于是将其下标放入栈中来更新**上一个没有被匹配的右括号的下标**。如果栈不为空，当前右括号的下标减去栈顶元素即为以该右括号为结尾的最长有效括号的长度。需要注意初始化时，不存在**上一个没有被匹配的右括号的下标**，那么将 `-1` 放入栈中，充当下标为 `0` 的“隔板”。时间复杂度 O(n)，空间复杂度 O(n)。
- 在栈的方法中，实际用到的元素仅仅是栈底的**上一个没有被匹配的右括号的下标**。那么考虑能否把这个值存在一个变量中，这样可以省去栈 O(n) 的时间复杂度。利用两个计数器 left 和 right 。首先，从左到右遍历字符串，每当遇到 `'('`，增加 left 计数器，每当遇到 `')'` ，增加 right 计数器。每当 left 计数器与 right 计数器相等时，计算当前有效字符串的长度，并且记录目前为止找到的最长子字符串。当 right 计数器比 left 计数器大时，说明括号不匹配，于是将 left 和 right 计数器同时变回 0。这样的做法利用了贪心的思想，考虑了以当前字符下标结尾的有效括号长度，每次当右括号数量多于左括号数量的时候之前的字符就扔掉不再考虑，重新从下一个字符开始计算。
- 但上面的做法会漏掉一种情况，就是遍历的时候左括号的数量始终大于右括号的数量，即 `(()` ，这种时候最长有效括号是求不出来的。解决办法是反向再计算一遍，如果从右往左计算，`(()` 先计算匹配的括号，最后只剩下 `'('`，这样依旧可以算出最长匹配的括号长度。反过来计算的方法和上述从左往右计算的方法一致：当 left 计数器比 right 计数器大时，将 left 和 right 计数器同时变回 0；当 left 计数器与 right 计数器相等时，计算当前有效字符串的长度，并且记录目前为止找到的最长子字符串。这种方法的时间复杂度是 O(n)，空间复杂度 O(1)。

## 代码

```go
package leetcode

// 解法一 栈
func longestValidParentheses(s string) int {
	stack, res := []int{}, 0
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
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

// 解法二 双指针
func longestValidParentheses1(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}
```