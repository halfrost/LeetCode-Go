# [515. Find Largest Value in Each Tree Row](https://leetcode.com/problems/find-largest-value-in-each-tree-row/)


## 题目

You need to find the largest value in each row of a binary tree.

**Example:**

    Input: 
    
              1
             / \
            3   2
           / \   \  
          5   3   9 
    
    Output: [1, 3, 9]


## 题目大意

求在二叉树的每一行中找到最大的值。


## 解题思路


- 给出一个二叉树，要求依次输出每行的最大值
- 用 BFS 层序遍历，将每层排序取出最大值。改进的做法是遍历中不断更新每层的最大值。

