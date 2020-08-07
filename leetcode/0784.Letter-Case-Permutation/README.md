# [784. Letter Case Permutation](https://leetcode.com/problems/letter-case-permutation/)


## 题目

Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string. Return a list of all possible strings we could create.

    Examples:
    Input: S = "a1b2"
    Output: ["a1b2", "a1B2", "A1b2", "A1B2"]
    
    Input: S = "3z4"
    Output: ["3z4", "3Z4"]
    
    Input: S = "12345"
    Output: ["12345"]

**Note:**

- `S` will be a string with length between `1` and `12`.
- `S` will consist only of letters or digits.


## 题目大意


给定一个字符串 S，通过将字符串 S 中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

## 解题思路


- 输出一个字符串中字母变大写，小写的所有组合。
- DFS 深搜或者 BFS 广搜都可以。
