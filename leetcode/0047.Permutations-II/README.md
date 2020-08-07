# [47. Permutations II](https://leetcode.com/problems/permutations-ii/)


## 题目

Given a collection of numbers that might contain duplicates, return all possible unique permutations.

**Example:**


    Input: [1,1,2]
    Output:
    [
      [1,1,2],
      [1,2,1],
      [2,1,1]
    ]


## 题目大意

给定一个可包含重复数字的序列，返回所有不重复的全排列。

## 解题思路

- 这一题是第 46 题的加强版，第 46 题中求数组的排列，数组中元素不重复，但是这一题中，数组元素会重复，所以需要最终排列出来的结果需要去重。
- 去重的方法是经典逻辑，将数组排序以后，判断重复元素再做逻辑判断。
- 其他思路和第 46 题完全一致，DFS 深搜即可。
