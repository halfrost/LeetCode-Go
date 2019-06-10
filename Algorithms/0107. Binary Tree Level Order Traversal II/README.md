# [107. Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/)

## 题目

Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],

```c
    3
   / \
  9  20
    /  \
   15   7
```

return its bottom-up level order traversal as:


```c
[
  [15,7],
  [9,20],
  [3]
]
```
 

## 题目大意

按层序从下到上遍历一颗树。

## 解题思路

用一个队列即可实现。



