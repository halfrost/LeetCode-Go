# [77. Combinations](https://leetcode.com/problems/combinations/)


## 题目

Given two integers *n* and *k*, return all possible combinations of *k* numbers out of 1 ... *n*.

**Example:**

    Input: n = 4, k = 2
    Output:
    [
      [2,4],
      [3,4],
      [2,3],
      [1,2],
      [1,3],
      [1,4],
    ]

## 题目大意

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

## 解题思路

- 计算排列组合中的组合，用 DFS 深搜即可，注意剪枝
