# [856. Score of Parentheses](https://leetcode.com/problems/score-of-parentheses/)

## 题目

Given a balanced parentheses string S, compute the score of the string based on the following rule:

() has score 1
AB has score A + B, where A and B are balanced parentheses strings.
(A) has score 2 * A, where A is a balanced parentheses string.
 

Example 1:

```c
Input: "()"
Output: 1

```

Example 2:

```c
Input: "(())"
Output: 2
```

Example 3:

```c
Input: "()()"
Output: 2
```

Example 4:

```c
Input: "(()(()))"
Output: 6
```
 

Note:

1. S is a balanced parentheses string, containing only ( and ).
2. 2 <= S.length <= 50

## 题目大意

按照以下规则计算括号的分数：() 代表 1 分。AB 代表 A + B，A 和 B 分别是已经满足匹配规则的括号组。(A) 代表 2 * A，其中 A 也是已经满足匹配规则的括号组。给出一个括号字符串，要求按照这些规则计算出括号的分数值。


## 解题思路

按照括号匹配的原则，一步步的计算每个组合的分数入栈。遇到题目中的 3 种情况，取出栈顶元素算分数。
