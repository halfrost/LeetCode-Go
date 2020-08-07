# [120. Triangle](https://leetcode.com/problems/triangle/)


## 题目

Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

    [
         [2],
        [3,4],
       [6,5,7],
      [4,1,8,3]
    ]

The minimum path sum from top to bottom is `11` (i.e., **2** + **3** + **5** + **1** = 11).

**Note:**

Bonus point if you are able to do this using only *O*(*n*) extra space, where *n* is the total number of rows in the triangle.


## 题目大意

给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。


## 解题思路

- 求出从三角形顶端到底端的最小和。要求最好用 O(n) 的时间复杂度。
- 这一题最优解是不用辅助空间，直接从下层往上层推。普通解法是用二维数组 DP，稍微优化的解法是一维数组 DP。解法如下：
