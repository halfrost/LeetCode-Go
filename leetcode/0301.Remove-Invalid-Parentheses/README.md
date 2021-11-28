# [301. Remove Invalid Parentheses](https://leetcode.com/problems/remove-invalid-parentheses/)


## 题目

Given a string s that contains parentheses and letters, remove the minimum number of invalid parentheses to make the input string valid.

Return all the possible results. You may return the answer in any order.

**Example 1:**

    Input: s = "()())()"
    Output: ["(())()","()()()"]

**Example 2:**

    Input: s = "(a)())()"
    Output: ["(a())()","(a)()()"]

**Example 3:**

    Input: s = ")("
    Output: [""]

**Constraints:**

- 1 <= s.length <= 25
- s consists of lowercase English letters and parentheses '(' and ')'.
- There will be at most 20 parentheses in s.

## 题目大意

给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

返回所有可能的结果。答案可以按 任意顺序 返回。

说明:

- 1 <= s.length <= 25
- s 由小写英文字母以及括号 '(' 和 ')' 组成
- s 中至多含 20 个括号


## 解题思路

回溯和剪枝
- 计算最大得分数maxScore，合法字符串的长度length，左括号和右括号的移除次数lmoves,rmoves
- 加一个左括号的得分加1；加一个右括号的得分减1
- 对于一个合法的字符串，左括号等于右括号，得分最终为0；
- 搜索过程中出现以下任何一种情况都直接返回
    - 得分值为负数
    - 得分大于最大得分数
    - 得分小于0
    - lmoves小于0
    - rmoves小于0

## 代码

```go

package leetcode

var (
	res      []string
	mp       map[string]int
	n        int
	length   int
	maxScore int
	str      string
)

func removeInvalidParentheses(s string) []string {
	lmoves, rmoves, lcnt, rcnt := 0, 0, 0, 0
	for _, v := range s {
		if v == '(' {
			lmoves++
			lcnt++
		} else if v == ')' {
			if lmoves != 0 {
				lmoves--
			} else {
				rmoves++
			}
			rcnt++
		}
	}
	n = len(s)
	length = n - lmoves - rmoves
	res = []string{}
	mp = make(map[string]int)
	maxScore = min(lcnt, rcnt)
	str = s
	backtrace(0, "", lmoves, rmoves, 0)
	return res
}

func backtrace(i int, cur string, lmoves int, rmoves int, score int) {
	if lmoves < 0 || rmoves < 0 || score < 0 || score > maxScore {
		return
	}
	if lmoves == 0 && rmoves == 0 {
		if len(cur) == length {
			if _, ok := mp[cur]; !ok {
				res = append(res, cur)
				mp[cur] = 1
			}
			return
		}
	}
	if i == n {
		return
	}
	if str[i] == '(' {
		backtrace(i+1, cur+string('('), lmoves, rmoves, score+1)
		backtrace(i+1, cur, lmoves-1, rmoves, score)
	} else if str[i] == ')' {
		backtrace(i+1, cur+string(')'), lmoves, rmoves, score-1)
		backtrace(i+1, cur, lmoves, rmoves-1, score)
	} else {
		backtrace(i+1, cur+string(str[i]), lmoves, rmoves, score)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```