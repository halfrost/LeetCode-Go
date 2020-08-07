# [303. Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/)


## 题目:

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

**Example:**

    Given nums = [-2, 0, 3, -5, 2, -1]
    
    sumRange(0, 2) -> 1
    sumRange(2, 5) -> -1
    sumRange(0, 5) -> -3

**Note:**

1. You may assume that the array does not change.
2. There are many calls to sumRange function.


## 题目大意

给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：

```
给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3

```

说明:

- 你可以假设数组不可变。
- 会多次调用 sumRange 方法。


## 解题思路


- 给出一个数组，数组里面的数都是`**不可变**`的，设计一个数据结构能够满足查询数组任意区间内元素的和。
- 这一题由于数组里面的元素都是`**不可变**`的，所以可以用 2 种方式来解答，第一种解法是用 prefixSum，通过累计和相减的办法来计算区间内的元素和，初始化的时间复杂度是 O(n)，但是查询区间元素和的时间复杂度是 O(1)。第二种解法是利用线段树，构建一颗线段树，父结点内存的是两个子结点的和，初始化建树的时间复杂度是 O(log n)，查询区间元素和的时间复杂度是 O(log n)。
