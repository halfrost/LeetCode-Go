# [700. Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/)

## 题目

You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

**Example 1**:

![https://assets.leetcode.com/uploads/2021/01/12/tree1.jpg](https://assets.leetcode.com/uploads/2021/01/12/tree1.jpg)

    Input: root = [4,2,7,1,3], val = 2
    Output: [2,1,3]

**Example 2**:

![https://assets.leetcode.com/uploads/2021/01/12/tree2.jpg](https://assets.leetcode.com/uploads/2021/01/12/tree2.jpg)

    Input: root = [4,2,7,1,3], val = 5
    Output: []

**Constraints:**

- The number of nodes in the tree is in the range [1, 5000].
- 1 <= Node.val <= 10000000
- root is a binary search tree.
- 1 <= val <= 10000000

## 题目大意

给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。

## 解题思路

- 根据二叉搜索树的性质(根节点的值大于左子树所有节点的值，小于右子树所有节点的值),进行递归求解

## 代码

```go

package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

```