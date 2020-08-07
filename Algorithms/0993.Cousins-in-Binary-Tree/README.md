# [993. Cousins in Binary Tree](https://leetcode.com/problems/cousins-in-binary-tree/)

## 题目

In a binary tree, the root node is at depth `0`, and children of each depth `k` node are at depth `k+1`.

Two nodes of a binary tree are *cousins* if they have the same depth, but have **different parents**.

We are given the `root` of a binary tree with unique values, and the values `x` and `y` of two different nodes in the tree.

Return `true` if and only if the nodes corresponding to the values `x` and `y` are cousins.

**Example 1:**

![](https://assets.leetcode.com/uploads/2019/02/12/q1248-01.png)

    Input: root = [1,2,3,4], x = 4, y = 3
    Output: false

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/02/12/q1248-02.png)

    Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
    Output: true

**Example 3:**

![](https://assets.leetcode.com/uploads/2019/02/13/q1248-03.png)

    Input: root = [1,2,3,null,4], x = 2, y = 3
    Output: false

**Note:**

1. The number of nodes in the tree will be between `2` and `100`.
2. Each node has a unique integer value from `1` to `100`.


## 题目大意

在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。如果二叉树的两个节点深度相同，但父节点不同，则它们是一对堂兄弟节点。我们给出了具有唯一值的二叉树的根节点 root，以及树中两个不同节点的值 x 和 y。只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true。否则，返回 false。



## 解题思路


- 给出一个二叉树，和 x ，y 两个值，要求判断这两个值是不是兄弟结点。兄弟结点的定义：都位于同一层，并且父结点是同一个结点。
- 这一题有 3 种解题方法，DFS、BFS、递归。思路都不难。
