# [216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/)


## 题目

Find all possible combinations of **k** numbers that add up to a number **n**, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.

**Note:**

- All numbers will be positive integers.
- The solution set must not contain duplicate combinations.

**Example 1:**

    Input: k = 3, n = 7
    Output: [[1,2,4]]

**Example 2:**

    Input: k = 3, n = 9
    Output: [[1,2,6], [1,3,5], [2,3,4]]

## 题目大意

找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

- 所有数字都是正整数。
- 解集不能包含重复的组合。


## 解题思路

- 这一题比第 39 题还要简单一些，在第 39 题上稍加改动就可以解出这一道题。
- 第 39 题是给出数组，这一道题数组是固定死的 [1,2,3,4,5,6,7,8,9]，并且数字不能重复使用。

