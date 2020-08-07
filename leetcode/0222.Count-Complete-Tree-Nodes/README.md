# [222. Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/)

## 题目


Given a complete binary tree, count the number of nodes.

Note:   

Definition of a complete binary tree from Wikipedia:
In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.


Example:

```c
Input: 
    1
   / \
  2   3
 / \  /
4  5 6

Output: 6
```

## 题目大意

输出一颗完全二叉树的结点个数。

## 解题思路

这道题其实按照层序遍历一次树，然后把每一层的结点个数相加即可。
