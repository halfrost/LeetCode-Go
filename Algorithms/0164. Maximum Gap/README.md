# [164. Maximum Gap](https://leetcode.com/problems/maximum-gap/)

## 题目

Given an unsorted array, find the maximum difference between the successive elements in its sorted form.

Return 0 if the array contains less than 2 elements.

Example 1:

```c
Input: [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either
             (3,6) or (6,9) has the maximum difference 3.
```

Example 2:

```c
Input: [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.
```


Note:

- You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
- Try to solve it in linear time/space.


## 题目大意

在数组中找到 2 个数字之间最大的间隔。要求尽量用 O(1) 的时间复杂度和空间复杂度。

## 解题思路

虽然使用排序算法可以 AC 这道题。先排序，然后依次计算数组中两两数字之间的间隔，找到最大的一个间隔输出即可。

这道题满足要求的做法是基数排序。




