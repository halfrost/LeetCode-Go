# [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/)

## 题目

Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example 1:

```c
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
```

Note:

- You must do this in-place without making a copy of the array.
- Minimize the total number of operations.



## 题目大意

题目要求不能采用额外的辅助空间，将数组中 0 元素都移动到数组的末尾，并且维持所有非 0 元素的相对位置。

## 解题思路

这一题可以只扫描数组一遍，不断的用 i，j 标记 0 和非 0 的元素，然后相互交换，最终到达题目的目的。与这一题相近的题目有第 26 题，第 27 题，第 80 题。
