# [301. Remove Invalid Parentheses](https://leetcode-cn.com/problems/remove-invalid-parentheses/)


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
