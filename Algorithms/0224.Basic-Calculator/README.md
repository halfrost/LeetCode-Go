# [224. Basic Calculator](https://leetcode.com/problems/basic-calculator/)


## 题目

Implement a basic calculator to evaluate a simple expression string.

The expression string may contain open `(` and closing parentheses `)`, the plus `+` or minus sign `-`, **non-negative** integers and empty spaces .

**Example 1:**

    Input: "1 + 1"
    Output: 2

**Example 2:**

    Input: " 2-1 + 2 "
    Output: 3

**Example 3:**

    Input: "(1+(4+5+2)-3)+(6+8)"
    Output: 23

**Note:**

- You may assume that the given expression is always valid.
- **Do not** use the `eval` built-in library function.

## 题目大意

实现一个基本的计算器来计算一个简单的字符串表达式的值。字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。

## 解题思路

- 注意点一：算式中有空格，需要跳过
- 注意点二：算式中会出现负数，负负得正的情况需要特殊处理，所以需要记录每次计算出来的符号

