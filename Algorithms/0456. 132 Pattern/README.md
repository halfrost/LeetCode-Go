# [456. 132 Pattern](https://leetcode.com/problems/132-pattern/)


## 题目

Given a sequence of n integers a1, a2, ..., an, a 132 pattern is a subsequence a**i**, a**j**, a**k** such that **i** < **j** < **k** and a**i** < a**k** < a**j**. Design an algorithm that takes a list of n numbers as input and checks whether there is a 132 pattern in the list.

**Note:** n will be less than 15,000.

**Example 1:**

    Input: [1, 2, 3, 4]
    
    Output: False
    
    Explanation: There is no 132 pattern in the sequence.

**Example 2:**

    Input: [3, 1, 4, 2]
    
    Output: True
    
    Explanation: There is a 132 pattern in the sequence: [1, 4, 2].

**Example 3:**

    Input: [-1, 3, 2, 0]
    
    Output: True
    
    Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].


## 题目大意


给定一个整数序列：a1, a2, ..., an，一个 132 模式的子序列 ai, aj, ak 被定义为：当 i < j < k 时，ai < ak < aj。设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有 132 模式的子序列。注意：n 的值小于 15000。


## 解题思路


- 这一题用暴力解法一定超时
- 这一题算是单调栈的经典解法，可以考虑从数组末尾开始往前扫，维护一个递减序列

