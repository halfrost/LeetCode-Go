# [52. N-Queens II](https://leetcode.com/problems/n-queens-ii/)


## 题目

The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

![](https://assets.leetcode.com/uploads/2018/10/12/8-queens.png)

Given an integer n, return the number of distinct solutions to the n-queens puzzle.

**Example:**


    Input: 4
    Output: 2
    Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
    [
     [".Q..",  // Solution 1
      "...Q",
      "Q...",
      "..Q."],
    
     ["..Q.",  // Solution 2
      "Q...",
      "...Q",
      ".Q.."]
    ]


## 题目大意

给定一个整数 n，返回 n 皇后不同的解决方案的数量。

## 解题思路

- 这一题是第 51 题的加强版，在第 51 题的基础上累加记录解的个数即可。
- 这一题也可以暴力打表法，时间复杂度为 O(1)。
