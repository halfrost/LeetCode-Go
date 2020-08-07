# [349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)

## 题目

Given two arrays, write a function to compute their intersection.


Example 1:

```c
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
```

Example 2:

```c
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
```

Note:

- Each element in the result must be unique.
- The result can be in any order.

## 题目大意

找到两个数组的交集元素，如果交集元素同一个数字出现了多次，只输出一次。

## 解题思路

把数组一的每个数字都存进字典中，然后在数组二中依次判断字典中是否存在，如果存在，在字典中删除它(因为输出要求只输出一次)。