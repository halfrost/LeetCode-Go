# [199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)

## 题目

Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example:

```c
Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
```

 

## 题目大意

从右边看一个树，输出看到的数字。注意有遮挡。


## 解题思路

- 这一题是按层序遍历的变种题。按照层序把每层的元素都遍历出来，然后依次取每一层的最右边的元素即可。用一个队列即可实现。
- 第 102 题和第 107 题都是按层序遍历的。


