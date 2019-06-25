# [39. Combination Sum](https://leetcode.com/problems/combination-sum/)


## 题目

Given a **set** of candidate numbers (`candidates`) **(without duplicates)** and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sums to `target`.

The **same** repeated number may be chosen from `candidates` unlimited number of times.

**Note:**

- All numbers (including `target`) will be positive integers.
- The solution set must not contain duplicate combinations.

**Example 1:**


    Input: candidates = [2,3,6,7], target = 7,
    A solution set is:
    [
      [7],
      [2,2,3]
    ]


**Example 2:**


    Input: candidates = [2,3,5], target = 8,
    A solution set is:
    [
      [2,2,2,2],
      [2,3,3],
      [3,5]
    ]


## 题目大意

给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。


## 解题思路

- 题目要求出总和为 sum 的所有组合，组合需要去重。
- 这一题和第 47 题类似，只不过元素可以反复使用。
