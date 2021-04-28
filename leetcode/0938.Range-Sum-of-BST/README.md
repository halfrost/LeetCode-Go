# [938. Range Sum of BST](https://leetcode.com/problems/range-sum-of-bst/)


## 题目

Given the `root` node of a binary search tree, return *the sum of values of all nodes with a value in the range `[low, high]`*.

**Example 1:**

![https://assets.leetcode.com/uploads/2020/11/05/bst1.jpg](https://assets.leetcode.com/uploads/2020/11/05/bst1.jpg)

```
Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
Output: 32

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/11/05/bst2.jpg](https://assets.leetcode.com/uploads/2020/11/05/bst2.jpg)

```
Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
Output: 23

```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 2 * 10^4]`.
- `1 <= Node.val <= 10^5`
- `1 <= low <= high <= 10^5`
- All `Node.val` are **unique**.

## 题目大意

给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。

## 解题思路

- 简单题。因为二叉搜索树的有序性，先序遍历即为有序。遍历过程中判断节点值是否位于区间范围内，在区间内就累加，不在区间内节点就不管。最终输出累加和。

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

func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	preOrder(root, low, high, &res)
	return res
}

func preOrder(root *TreeNode, low, high int, res *int) {
	if root == nil {
		return
	}
	if low <= root.Val && root.Val <= high {
		*res += root.Val
	}
	preOrder(root.Left, low, high, res)
	preOrder(root.Right, low, high, res)
}
```