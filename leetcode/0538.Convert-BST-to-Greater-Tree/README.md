# [538. Convert BST to Greater Tree](https://leetcode.com/problems/convert-bst-to-greater-tree/)

## 题目

Given the `root` of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

As a reminder, a *binary search tree* is a tree that satisfies these constraints:

- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

**Note:** This question is the same as 1038: [https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/](https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/)

**Example 1:**

![https://assets.leetcode.com/uploads/2019/05/02/tree.png](https://assets.leetcode.com/uploads/2019/05/02/tree.png)

```
Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
```

**Example 2:**

```
Input: root = [0,null,1]
Output: [1,null,1]
```

**Example 3:**

```
Input: root = [1,0,2]
Output: [3,3,2]
```

**Example 4:**

```
Input: root = [3,2,4,1]
Output: [7,9,4,10]
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 104]`.
- `104 <= Node.val <= 104`
- All the values in the tree are **unique**.
- `root` is guaranteed to be a valid binary search tree.

## 题目大意

给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

- 节点的左子树仅包含键 小于 节点键的节点。
- 节点的右子树仅包含键 大于 节点键的节点。
- 左右子树也必须是二叉搜索树。

## 解题思路

- 根据二叉搜索树的有序性，想要将其转换为累加树，只需按照 右节点 - 根节点 - 左节点的顺序遍历，并累加和即可。
- 此题同第 1038 题。

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

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	dfs538(root, &sum)
	return root
}

func dfs538(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	dfs538(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	dfs538(root.Left, sum)
}
```