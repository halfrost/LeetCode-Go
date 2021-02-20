# [1249. Minimum Remove to Make Valid Parentheses](https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/)


## 题目

Given a string s of `'('` , `')'` and lowercase English characters.

Your task is to remove the minimum number of parentheses ( `'('` or `')'`, in any positions ) so that the resulting *parentheses string* is valid and return **any** valid string.

Formally, a *parentheses string* is valid if and only if:

- It is the empty string, contains only lowercase characters, or
- It can be written as `AB` (`A` concatenated with `B`), where `A` and `B` are valid strings, or
- It can be written as `(A)`, where `A` is a valid string.

**Example 1:**

```
Input: s = "lee(t(c)o)de)"
Output: "lee(t(c)o)de"
Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.

```

**Example 2:**

```
Input: s = "a)b(c)d"
Output: "ab(c)d"

```

**Example 3:**

```
Input: s = "))(("
Output: ""
Explanation: An empty string is also valid.

```

**Example 4:**

```
Input: s = "(a(b(c)d)"
Output: "a(b(c)d)"

```

**Constraints:**

- `1 <= s.length <= 10^5`
- `s[i]` is one of `'('` , `')'` and lowercase English letters`.`

## 题目大意

给你一个由 '('、')' 和小写字母组成的字符串 s。你需要从字符串中删除最少数目的 '(' 或者 ')' （可以删除任意位置的括号)，使得剩下的「括号字符串」有效。请返回任意一个合法字符串。有效「括号字符串」应当符合以下 任意一条 要求：

- 空字符串或只包含小写字母的字符串
- 可以被写作 AB（A 连接 B）的字符串，其中 A 和 B 都是有效「括号字符串」
- 可以被写作 (A) 的字符串，其中 A 是一个有效的「括号字符串」

## 解题思路

- 最容易想到的思路是利用栈判断括号匹配是否有效。这个思路可行，时间复杂度也只是 O(n)。
- 不用栈，可以 2 次循环遍历，正向遍历一次，标记出多余的 `'('` ，逆向遍历一次，再标记出多余的 `')'`，最后将所有这些标记多余的字符删掉即可。这种解法写出来的代码也很简洁，时间复杂度也是 O(n)。
- 针对上面的解法再改进一点。正向遍历的时候不仅标记出多余的 `'('`，还可以顺手把多余的 `')'` 删除。这样只用循环一次。最后再删除掉多余的 `'('` 即可。时间复杂度还是 O(n)。

## 代码

```go
package leetcode

func minRemoveToMakeValid(s string) string {
	res, opens := []byte{}, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			opens++
		} else if s[i] == ')' {
			if opens == 0 {
				continue
			}
			opens--
		}
		res = append(res, s[i])
	}
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == '(' && opens > 0 {
			opens--
			res = append(res[:i], res[i+1:]...)
		}
	}
	return string(res)
}
```