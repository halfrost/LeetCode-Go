# [131. Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/)


## 题目

Given a string *s*, partition *s* such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of *s*.

**Example:**

    Input: "aab"
    Output:
    [
      ["aa","b"],
      ["a","a","b"]
    ]

## 题目大意

给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。返回 s 所有可能的分割方案。

## 解题思路

- 要求输出一个字符串可以被拆成回文串的所有解，DFS 递归求解即可。
