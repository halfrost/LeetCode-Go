# [719. Find K-th Smallest Pair Distance](https://leetcode.com/problems/find-k-th-smallest-pair-distance/)


## 题目:

Given an integer array, return the k-th smallest **distance** among all the pairs. The distance of a pair (A, B) is defined as the absolute difference between A and B.

**Example 1:**

    Input:
    nums = [1,3,1]
    k = 1
    Output: 0 
    Explanation:
    Here are all the pairs:
    (1,3) -> 2
    (1,1) -> 0
    (3,1) -> 2
    Then the 1st smallest distance pair is (1,1), and its distance is 0.

**Note:**

1. `2 <= len(nums) <= 10000`.
2. `0 <= nums[i] < 1000000`.
3. `1 <= k <= len(nums) * (len(nums) - 1) / 2`.


## 题目大意

给定一个整数数组，返回所有数对之间的第 k 个最小距离。一对 (A, B) 的距离被定义为 A 和 B 之间的绝对差值。

提示:

1. 2 <= len(nums) <= 10000.
2. 0 <= nums[i] < 1000000.
3. 1 <= k <= len(nums) * (len(nums) - 1) / 2.



## 解题思路

- 给出一个数组，要求找出第 k 小两两元素之差的值。两两元素之差可能重复，重复的元素之差算多个，不去重。
- 这一题可以用二分搜索来解答。先把原数组排序，那么最大的差值就是 `nums[len(nums)-1] - nums[0]` ，最小的差值是 0，即在 `[0, nums[len(nums)-1] - nums[0]]` 区间内搜索最终答案。针对每个 `mid`，判断小于等于 `mid` 的差值有多少个。题意就转化为，在数组中找到这样一个数，使得满足 `nums[i] - nums[j] ≤ mid` 条件的组合数等于 `k`。那么如何计算满足两两数的差值小于 mid 的组合总数是本题的关键。
- 最暴力的方法就是 2 重循环，暴力计数。这个方法效率不高，耗时很长。原因是没有利用数组有序这一条件。实际上数组有序对计算满足条件的组合数有帮助。利用双指针滑动即可计算出组合总数。见解法一。
