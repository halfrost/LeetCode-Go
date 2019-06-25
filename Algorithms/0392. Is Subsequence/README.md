# [392. Is Subsequence](https://leetcode.com/problems/is-subsequence/)


## 题目

Given a string **s** and a string **t**, check if **s** is subsequence of **t**.

You may assume that there is only lower case English letters in both **s** and **t**. **t** is potentially a very long (length ~= 500,000) string, and **s** is a short string (<=100).

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, `"ace"` is a subsequence of `"abcde"`while `"aec"` is not).

**Example 1:**

**s** = `"abc"`, **t** = `"ahbgdc"`

Return `true`.

**Example 2:**

**s** = `"axc"`, **t** = `"ahbgdc"`

Return `false`.

**Follow up:**If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you want to check one by one to see if T has its subsequence. In this scenario, how would you change your code?

**Credits:**Special thanks to [@pbrother](https://leetcode.com/pbrother/) for adding this problem and creating all test cases.


## 题目大意

给定字符串 s 和 t ，判断 s 是否为 t 的子序列。你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。



## 解题思路


- 给定 2 个字符串 s 和 t，问 s 是不是 t 的子序列。注意 s 在 t 中还需要保持 s 的字母的顺序。
- 这是一题贪心算法。直接做即可。

