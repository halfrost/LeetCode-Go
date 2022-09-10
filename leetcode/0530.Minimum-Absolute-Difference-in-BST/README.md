# [530. Minimum Absolute Difference in BST](https://leetcode.com/problems/minimum-absolute-difference-in-bst/)


## 题目

Given a binary search tree with non-negative values, find the minimum [absolute difference](https://en.wikipedia.org/wiki/Absolute_difference) between values of any two nodes.

**Example:**

```
Input:

   1
    \
     3
    /
   2

Output:
1

Explanation:
The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).
```

**Note:**

- There are at least two nodes in this BST.
- This question is the same as 783: [https://leetcode.com/problems/minimum-distance-between-bst-nodes/](https://leetcode.com/problems/minimum-distance-between-bst-nodes/)

## 题目大意

给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

## 解题思路

- 由于是 BST 树，利用它有序的性质，中根遍历的结果是有序的。中根遍历过程中动态维护前后两个节点的差值，即可找到最小差值。
- 此题与第 783 题完全相同。

## 代码

```go
package leetcode

import (
	"math"

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

func getMinimumDifference(root *TreeNode) int {
	res, nodes := math.MaxInt16, -1
	dfsBST(root, &res, &nodes)
	return res
}

func dfsBST(root *TreeNode, res, pre *int) {
	if root == nil {
		return
	}
	dfsBST(root.Left, res, pre)
	if *pre != -1 {
		*res = min(*res, abs(root.Val-*pre))
	}
	*pre = root.Val
	dfsBST(root.Right, res, pre)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```