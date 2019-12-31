# [513. Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value/)


## 题目:

Given a binary tree, find the leftmost value in the last row of the tree.

**Example 1:**

    Input:
    
        2
       / \
      1   3
    
    Output:
    1

**Example 2:**

    Input:
    
            1
           / \
          2   3
         /   / \
        4   5   6
           /
          7
    
    Output:
    7

**Note:** You may assume the tree (i.e., the given root node) is not **NULL**.


## 题目大意

给定一个二叉树，在树的最后一行找到最左边的值。注意: 您可以假设树（即给定的根节点）不为 NULL。






## 解题思路


- 给出一棵树，输出这棵树最下一层中最左边的节点的值。
- 这一题用 DFS 和 BFS 均可解题。
