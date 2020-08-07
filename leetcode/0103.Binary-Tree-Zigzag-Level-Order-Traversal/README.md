# [103. Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)

## 题目

Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],

```c
    3
   / \
  9  20
    /  \
   15   7
```

return its zigzag level order traversal as:

```c
[
  [3],
  [20,9],
  [15,7]
]
```
 

## 题目大意

按照 Z 字型层序遍历一棵树。

## 解题思路

- 按层序从上到下遍历一颗树，但是每一层的顺序是相互反转的，即上一层是从左往右，下一层就是从右往左，以此类推。用一个队列即可实现。
- 第 102 题和第 107 题都是按层序遍历的。


