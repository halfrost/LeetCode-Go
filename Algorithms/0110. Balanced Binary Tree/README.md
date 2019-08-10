# [110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)

## 题目


Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example 1:

Given the following tree [3,9,20,null,null,15,7]:

```c
    3
   / \
  9  20
    /  \
   15   7
```

Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:


```c
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
```

Return false.


## 题目大意

判断一棵树是不是平衡二叉树。平衡二叉树的定义是：树中每个节点都满足左右两个子树的高度差 <= 1 的这个条件。


## 解题思路

根据定义判断即可，计算树的高度是第 104 题。


