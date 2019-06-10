# [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

## 题目

Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

```c
Input: [3,2,1,5,6,4] and k = 2
Output: 5
```

Example 2:

```c
Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
```

Note:     

You may assume k is always valid, 1 ≤ k ≤ array's length.


## 题目大意

找出数组中第 K 大的元素。这一题非常经典。可以用 O(n) 的时间复杂度实现。

## 解题思路

在快排的 partition 操作中，每次 partition 操作结束都会返回一个点，这个标定点的下标和最终排序之后有序数组中这个元素所在的下标是一致的。利用这个特性，我们可以不断的划分数组区间，最终找到第 K 大的元素。执行一次 partition 操作以后，如果这个元素的下标比 K 小，那么接着就在后边的区间继续执行 partition 操作；如果这个元素的下标比 K 大，那么就在左边的区间继续执行 partition 操作；如果相等就直接输出这个下标对应的数组元素即可。

