# [229. Majority Element II](https://leetcode.com/problems/majority-element-ii/)


## 题目:

Given an integer array of size n, find all elements that appear more than `⌊ n/3 ⌋` times.

**Note:** The algorithm should run in linear time and in O(1) space.

**Example 1:**

    Input: [3,2,3]
    Output: [3]

**Example 2:**

    Input: [1,1,1,3,3,2,2,2]
    Output: [1,2]


## 题目大意

给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。

## 解题思路

- 这一题是第 169 题的加强版。Boyer-Moore Majority Vote algorithm 算法的扩展版。
- 题目要求找出数组中出现次数大于 `⌊ n/3 ⌋` 次的数。要求空间复杂度为 O(1)。简单题。
- 这篇文章写的不错，可参考：[https://gregable.com/2013/10/majority-vote-algorithm-find-majority.html](https://gregable.com/2013/10/majority-vote-algorithm-find-majority.html)