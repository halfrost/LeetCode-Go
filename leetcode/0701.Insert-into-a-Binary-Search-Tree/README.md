# [701. Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/)


## 题目

You are given the `root` node of a binary search tree (BST) and a `value` to insert into the tree. Return *the root node of the BST after the insertion*. It is **guaranteed** that the new value does not exist in the original BST.

**Notice** that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return **any of them**.

**Example 1:**

![https://assets.leetcode.com/uploads/2020/10/05/insertbst.jpg](https://assets.leetcode.com/uploads/2020/10/05/insertbst.jpg)

```
Input: root = [4,2,7,1,3], val = 5
Output: [4,2,7,1,3,5]
Explanation: Another accepted tree is:

```

![https://assets.leetcode.com/uploads/2020/10/05/bst.jpg](https://assets.leetcode.com/uploads/2020/10/05/bst.jpg)

**Example 2:**

```
Input: root = [40,20,60,10,30,50,70], val = 25
Output: [40,20,60,10,30,50,70,null,null,25]

```

**Example 3:**

```
Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
Output: [4,2,7,1,3,5]

```

**Constraints:**

- The number of nodes in the tree will be in the range `[0, 104]`.
- `108 <= Node.val <= 108`
- All the values `Node.val` are **unique**.
- `108 <= val <= 108`
- It's **guaranteed** that `val` does not exist in the original BST.

## 题目大意

给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。

## 解题思路

- 简单题。插入节点的方法有多种，笔者这里选择一种简单的方法。从根开始遍历这个二叉树，当前节点的值比待插入节点的值小，则往右遍历；当前节点的值比待插入节点的值大，则往左遍历。最后遍历到空节点便是要插入的地方。

## 代码

```go
package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

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
func insert(n *TreeNode, val int) *TreeNode {
	if n == nil {
		return &TreeNode{Val: val}
	}
	if n.Val < val {
		n.Right = insert(n.Right, val)
	} else {
		n.Left = insert(n.Left, val)
	}
	return n
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return insert(root, val)
}
```