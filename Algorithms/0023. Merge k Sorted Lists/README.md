# [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)

## 题目

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.



Example :

```
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6

```

## 题目大意

合并 K 个有序链表

## 解题思路

借助分治的思想，把 K 个有序链表两两合并即可。相当于是第 21 题的加强版。