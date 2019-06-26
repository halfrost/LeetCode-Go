# [59. Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/)


## 题目

Given a positive integer *n*, generate a square matrix filled with elements from 1 to *n*2 in spiral order.

**Example:**


    Input: 3
    Output:
    [
     [ 1, 2, 3 ],
     [ 8, 9, 4 ],
     [ 7, 6, 5 ]
    ]


## 题目大意

给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。


## 解题思路

- 给出一个数组 n，要求输出一个 n * n 的二维数组，里面元素是 1 - n*n，且数组排列顺序是螺旋排列的
- 这一题是第 54 题的加强版，没有需要注意的特殊情况，直接模拟即可。

