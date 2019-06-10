# [977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/)

## 题目

Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.



Example 1:

```c
Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
```

Example 2:

```c
Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]
```

Note:

1. 1 <= A.length <= 10000
2. -10000 <= A[i] <= 10000
3. A is sorted in non-decreasing order.

## 题目大意

给一个已经有序的数组，返回的数组也必须是有序的，且数组中的每个元素是由原数组中每个数字的平方得到的。

## 解题思路

这一题由于原数组是有序的，所以要尽量利用这一特点来减少时间复杂度。

最终返回的数组，最后一位，是最大值，这个值应该是由原数组最大值，或者最小值得来的，所以可以从数组的最后一位开始排列最终数组。用 2 个指针分别指向原数组的首尾，分别计算平方值，然后比较两者大小，大的放在最终数组的后面。然后大的一个指针移动。直至两个指针相撞，最终数组就排列完成了。










