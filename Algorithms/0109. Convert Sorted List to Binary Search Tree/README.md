# [109. Convert Sorted List to Binary Search Tree](https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/)

## 题目

Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

```
Given the sorted linked list: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
```


## 题目大意

将链表转化为高度平衡的二叉搜索树。高度平衡的定义：每个结点的 2 个子结点的深度不能相差超过 1 。

## 解题思路

思路比较简单，依次把链表的中间点作为根结点，类似二分的思想，递归排列所有结点即可。