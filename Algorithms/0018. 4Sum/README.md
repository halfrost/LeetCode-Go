# [18. 4Sum](https://leetcode.com/problems/4sum/)

## 题目

Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

```c
Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
```

## 题目大意

给定一个数组，要求在这个数组中找出 4 个数之和为 0 的所有组合。


## 解题思路

用 map 提前计算好任意 3 个数字之和，保存起来，可以将时间复杂度降到 O(n^3)。这一题比较麻烦的一点在于，最后输出解的时候，要求输出不重复的解。数组中同一个数字可能出现多次，同一个数字也可能使用多次，但是最后输出解的时候，不能重复。例如 [-1，1，2, -2] 和 [2, -1, -2, 1]、[-2, 2, -1, 1] 这 3 个解是重复的，即使 -1, -2 可能出现 100 次，每次使用的 -1, -2 的数组下标都是不同的。

这一题是第 15 题的升级版，思路都是完全一致的。这里就需要去重和排序了。map 记录每个数字出现的次数，然后对 map 的 key 数组进行排序，最后在这个排序以后的数组里面扫，找到另外 3 个数字能和自己组成 0 的组合。

第 15 题和第 18 题的解法一致。



