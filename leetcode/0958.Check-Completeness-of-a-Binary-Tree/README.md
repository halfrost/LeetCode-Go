# [958. Check Completeness of a Binary Tree](https://leetcode.com/problems/check-completeness-of-a-binary-tree/)

## 题目

Given the root of a binary tree, determine if it is a complete binary tree.

In a complete binary tree, every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Example 1:
```c
Input: root = [1,2,3,4,5,6]
Output: true
Explanation: Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.

     1
   /   \
  2     3 
 / \   /
4   5 6    
```
Example 2:
```c
Input: root = [1,2,3,4,5,null,7]
Output: false
Explanation: The node with value 7 isn't as far left as possible.

     1
   /   \
  2     3 
 / \     \
4   5     7
```

Constraints:

The number of nodes in the tree is in the range [1, 100].
1 <= Node.val <= 1000


## 题目大意

若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）

说明要判断每个节点的左孩子必须不为空。


## 解题思路

- 这一题是按层序遍历的变种题。
- 判断每个节点的左孩子是否为空。
- 第 102,107,199 都是按层序遍历的。


