# [996. Number of Squareful Arrays](https://leetcode.com/problems/number-of-squareful-arrays/)



## 题目

Given an array `A` of non-negative integers, the array is *squareful* if for every pair of adjacent elements, their sum is a perfect square.

Return the number of permutations of A that are squareful. Two permutations `A1` and `A2` differ if and only if there is some index `i` such that `A1[i] != A2[i]`.

**Example 1:**

    Input: [1,17,8]
    Output: 2
    Explanation: 
    [1,8,17] and [17,8,1] are the valid permutations.

**Example 2:**

    Input: [2,2,2]
    Output: 1

**Note:**

1. `1 <= A.length <= 12`
2. `0 <= A[i] <= 1e9`


## 题目大意

给定一个非负整数数组 A，如果该数组每对相邻元素之和是一个完全平方数，则称这一数组为正方形数组。

返回 A 的正方形排列的数目。两个排列 A1 和 A2 不同的充要条件是存在某个索引 i，使得 A1[i] != A2[i]。



## 解题思路


- 这一题是第 47 题的加强版。第 47 题要求求出一个数组的所有不重复的排列。这一题要求求出一个数组的所有不重复，且相邻两个数字之和都为完全平方数的排列。
- 思路和第 47 题完全一致，只不过增加判断相邻两个数字之和为完全平方数的判断，注意在 DFS 的过程中，需要剪枝，否则时间复杂度很高，会超时。
