# [167. Two Sum II - Input array is sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)

## 题目

Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

- Your returned answers (both index1 and index2) are not zero-based.
- You may assume that each input would have exactly one solution and you may not use the same element twice.

Example:

```c
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
```

## 题目大意

找出两个数之和等于 target 的两个数字，要求输出它们的下标。注意一个数字不能使用 2 次。下标从小到大输出。假定题目一定有一个解。

## 解题思路

这一题比第 1 题 Two Sum 的问题还要简单，因为这里数组是有序的。可以直接用第一题的解法解决这道题。

