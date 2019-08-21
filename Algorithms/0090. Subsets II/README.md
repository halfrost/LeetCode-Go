# [90. Subsets II](https://leetcode.com/problems/subsets-ii/)


## 题目

Given a collection of integers that might contain duplicates, ***nums***, return all possible subsets (the power set).

**Note:** The solution set must not contain duplicate subsets.

**Example:**

    Input: [1,2,2]
    Output:
    [
      [2],
      [1],
      [1,2,2],
      [2,2],
      [1,2],
      []
    ]

## 题目大意

给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。说明：解集不能包含重复的子集。


## 解题思路

- 这一题是第 78 题的加强版，比第 78 题多了一个条件，数组中的数字会出现重复。
- 解题方法依旧是 DFS，需要在回溯的过程中加上一些判断。
- 这一题和第 78 题，第 491 题类似，可以一起解答和复习。

