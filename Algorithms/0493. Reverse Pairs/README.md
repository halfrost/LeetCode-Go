# [493. Reverse Pairs](https://leetcode.com/problems/reverse-pairs/)


## 题目:

Given an array `nums`, we call `(i, j)` an **important reverse pair** if `i < j` and `nums[i] > 2*nums[j]`.

You need to return the number of important reverse pairs in the given array.

**Example1:**

    Input: [1,3,2,3,1]
    Output: 2

**Example2:**

    Input: [2,4,3,5,1]
    Output: 3

**Note:**

1. The length of the given array will not exceed `50,000`.
2. All the numbers in the input array are in the range of 32-bit integer.


## 题目大意

给定一个数组 nums ，如果 i < j 且 nums[i] > 2\*nums[j] 我们就将 (i, j) 称作一个重要翻转对。你需要返回给定数组中的重要翻转对的数量。

注意:

- 给定数组的长度不会超过 50000。
- 输入数组中的所有数字都在 32 位整数的表示范围内。


## 解题思路


- 给出一个数组，要求找出满足条件的所有的“重要的反转对” (i,j)。重要的反转对的定义是：`i<j`，并且 `nums[i] > 2*nums[j]`。
- 这一题是 327 题的变种题。首先将数组中所有的元素以及各自的 `2*nums[i] + 1` 都放在字典中去重。去重以后再做离散化处理。这一题的测试用例会卡离散化，如果不离散化，Math.MaxInt32 会导致数字溢出，见测试用例中 2147483647, -2147483647 这组测试用例。离散后，映射关系 保存在字典中。从左往右遍历数组，先 query ，再 update ，这个顺序和第 327 题是反的。先 query 查找 `[2*nums[i] + 1, len(indexMap)-1]` 这个区间内满足条件的值，这个区间内的值都是 `> 2*nums[j]` 的。这一题移动的是 `j`，`j` 不断的变化，往线段树中不断插入的是 `i`。每轮循环先 query 一次前一轮循环中累积插入线段树中的 `i`，这些累积在线段树中的代表的是所有在 `j` 前面的 `i`。query 查询的是本轮 `[2*nums[j] + 1, len(indexMap)-1]`，如果能找到，即找到了这样一个 `j`，能满足 `nums[i] > 2*nums[j`， 把整个数组都扫完，累加的 query 出来的 count 计数就是最终答案。
- 类似的题目：第 327 题，第 315 题。
- 这一题用线段树并不是最优解，用线段树解这一题是为了训练线段树这个数据结构。最优解是解法二中的 mergesort。
