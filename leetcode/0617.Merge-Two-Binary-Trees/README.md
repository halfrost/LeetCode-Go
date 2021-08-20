# [617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/)


## 题目

You are given two binary trees `root1` and `root2`.

Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.

Return *the merged tree*.

**Note:** The merging process must start from the root nodes of both trees.

**Example 1:**

![https://assets.leetcode.com/uploads/2021/02/05/merge.jpg](https://assets.leetcode.com/uploads/2021/02/05/merge.jpg)

```
Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
Output: [3,4,5,5,4,null,7]
```

**Example 2:**

```
Input: root1 = [1], root2 = [1,2]
Output: [2,2]
```

**Constraints:**

- The number of nodes in both trees is in the range `[0, 2000]`.
- `104 <= Node.val <= 104`

## 题目大意

给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

## 解题思路

- 简单题。采用深搜的思路，分别从根节点开始同时遍历两个二叉树，并将对应的节点进行合并。两个二叉树的对应节点可能存在以下三种情况：
    - 如果两个二叉树的对应节点都为空，则合并后的二叉树的对应节点也为空；
    - 如果两个二叉树的对应节点只有一个为空，则合并后的二叉树的对应节点为其中的非空节点；
    - 如果两个二叉树的对应节点都不为空，则合并后的二叉树的对应节点的值为两个二叉树的对应节点的值之和，此时需要显性合并两个节点。
- 对一个节点进行合并之后，还要对该节点的左右子树分别进行合并。用递归实现即可。

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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
```