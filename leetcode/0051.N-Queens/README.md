# [51. N-Queens](https://leetcode.com/problems/n-queens/)


## 题目

The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

![](https://assets.leetcode.com/uploads/2018/10/12/8-queens.png)

Given an integer *n*, return all distinct solutions to the *n*-queens puzzle.

Each solution contains a distinct board configuration of the *n*-queens' placement, where `'Q'` and `'.'` both indicate a queen and an empty space respectively.

**Example:**


    Input: 4
    Output: [
     [".Q..",  // Solution 1
      "...Q",
      "Q...",
      "..Q."],
    
     ["..Q.",  // Solution 2
      "Q...",
      "...Q",
      ".Q.."]
    ]
    Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.


## 题目大意

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。


## 解题思路

- 求解 n 皇后问题
- 利用 col 数组记录列信息，col 有 `n` 列。用 dia1，dia2 记录从左下到右上的对角线，从左上到右下的对角线的信息，dia1 和 dia2 分别都有 `2*n-1` 个。
- dia1 对角线的规律是 `i + j 是定值`，例如[0,0]，为 0；[1,0]、[0,1] 为 1；[2,0]、[1,1]、[0,2] 为 2；
- dia2 对角线的规律是 `i - j 是定值`，例如[0,7]，为 -7；[0,6]、[1,7] 为 -6；[0,5]、[1,6]、[2,7] 为 -5；为了使他们从 0 开始，i - j + n - 1 偏移到 0 开始，所以 dia2 的规律是 `i - j + n - 1 为定值`。

