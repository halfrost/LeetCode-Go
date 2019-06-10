# [287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)

## 题目

Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Example 1:

```c
Input: [1,3,4,2,2]
Output: 2
```

Example 2:

```c
Input: [3,1,3,4,2]
Output: 3
```

Note:  

- You must not modify the array (assume the array is read only).
- You must use only constant, O(1) extra space.
- Your runtime complexity should be less than O(n2).
- There is only one duplicate number in the array, but it could be repeated more than once.

## 题目大意

给出 n + 1 个数，这些数是在 1-n 中取值的，同一个数字可以出现多次。要求找出这些数中重复的数字。时间复杂度最好低于 O(n^2)，空间复杂度为 O(1)。

## 解题思路

这道题比较巧的思路是，将这些数字想象成链表中的结点，数组中数字代表下一个结点的数组下标。找重复的数字就是找链表中成环的那个点。由于题目保证了一定会有重复的数字，所以一定会成环。所以用快慢指针的方法，快指针一次走 2 步，慢指针一次走 1 步，相交以后，快指针从头开始，每次走一步，再次遇见的时候就是成环的交点处，也即是重复数字所在的地方。


另外一个做法是，先将数组排序，依照下标是从 0 开始递增的特性，那么数组里面的数字与下标的差值应该是越来越大。如果出现了相同的数字，下标变大，差值应该比前一个数字小，出现了这个情况就说明找到了相同数字了。