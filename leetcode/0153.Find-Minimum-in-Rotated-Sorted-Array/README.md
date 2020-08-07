# [153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)


## 题目:

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., `[0,1,2,4,5,6,7]` might become `[4,5,6,7,0,1,2]`).

Find the minimum element.

You may assume no duplicate exists in the array.

**Example 1:**

    Input: [3,4,5,1,2] 
    Output: 1

**Example 2:**

    Input: [4,5,6,7,0,1,2]
    Output: 0


## 题目大意

假设按照升序排序的数组在预先未知的某个点上进行了旋转。( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。请找出其中最小的元素。

你可以假设数组中不存在重复元素。


## 解题思路

- 给出一个原本从小到大排序过的数组，但是在某一个分割点上，把数组切分后的两部分对调位置，数值偏大的放到了数组的前部。求这个数组中最小的元素。
- 求数组最小的元素其实就是找分割点，前一个数比当前数大，后一个数比当前数也要大。可以用二分搜索查找，需要查找的两个有序区间。时间复杂度  O(log n)。这一题也可以用暴力解法，从头开始遍历，动态维护一个最小值即可，时间复杂度 O(n)。
