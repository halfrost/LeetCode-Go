# [220. Contains Duplicate III](https://leetcode.com/problems/contains-duplicate-iii/)

## 题目

Given an array of integers, find out whether there are two distinct indices i and j in the array such that the absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.

Example 1:

```c
Input: nums = [1,2,3,1], k = 3, t = 0
Output: true
```

Example 2:

```c
Input: nums = [1,0,1,1], k = 1, t = 2
Output: true
```
Example 3:

```c
Input: nums = [1,5,9,1,5,9], k = 2, t = 3
Output: false
```

## 题目大意

给出一个数组 num，再给 K 和 t。问在 num 中能否找到一组 i 和 j，使得 num[i] 和 num[j] 的绝对差值最大为 t，并且 i 和 j 之前的绝对差值最大为 k。

## 解题思路

这道题就按照题意来，两层 for 循环。