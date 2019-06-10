# [202. Happy Number](https://leetcode.com/problems/happy-number/)

## 题目

Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Example 1:

```c
Input: 19
Output: true
Explanation: 
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
```

## 题目大意

判断一个数字是否是“快乐数字”，“快乐数字”的定义是，不断的把这个数字的每个数字的平方和加起来，反复的加，最终如果能有结果是 1，则是“快乐数字”，如果不能得到一，出现了循环，则输出 false。

## 解题思路

按照题意要求做即可。

