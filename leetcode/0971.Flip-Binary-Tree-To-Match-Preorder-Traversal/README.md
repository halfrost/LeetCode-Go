# [971. Flip Binary Tree To Match Preorder Traversal](https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/)


## 题目

You are given the `root` of a binary tree with `n` nodes, where each node is uniquely assigned a value from `1` to `n`. You are also given a sequence of `n` values `voyage`, which is the **desired** **[pre-order traversal](https://en.wikipedia.org/wiki/Tree_traversal#Pre-order)** of the binary tree.

Any node in the binary tree can be **flipped** by swapping its left and right subtrees. For example, flipping node 1 will have the following effect:

![https://assets.leetcode.com/uploads/2021/02/15/fliptree.jpg](https://assets.leetcode.com/uploads/2021/02/15/fliptree.jpg)

Flip the **smallest** number of nodes so that the **pre-order traversal** of the tree **matches** `voyage`.

Return *a list of the values of all **flipped** nodes. You may return the answer in **any order**. If it is **impossible** to flip the nodes in the tree to make the pre-order traversal match* `voyage`*, return the list* `[-1]`.

**Example 1:**

![https://assets.leetcode.com/uploads/2019/01/02/1219-01.png](https://assets.leetcode.com/uploads/2019/01/02/1219-01.png)

```
Input: root = [1,2], voyage = [2,1]
Output: [-1]
Explanation: It is impossible to flip the nodes such that the pre-order traversal matches voyage.
```

**Example 2:**

![https://assets.leetcode.com/uploads/2019/01/02/1219-02.png](https://assets.leetcode.com/uploads/2019/01/02/1219-02.png)

```
Input: root = [1,2,3], voyage = [1,3,2]
Output: [1]
Explanation: Flipping node 1 swaps nodes 2 and 3, so the pre-order traversal matches voyage.
```

**Example 3:**

![https://assets.leetcode.com/uploads/2019/01/02/1219-02.png](https://assets.leetcode.com/uploads/2019/01/02/1219-02.png)

```
Input: root = [1,2,3], voyage = [1,2,3]
Output: []
Explanation: The tree's pre-order traversal already matches voyage, so no nodes need to be flipped.
```

**Constraints:**

- The number of nodes in the tree is `n`.
- `n == voyage.length`
- `1 <= n <= 100`
- `1 <= Node.val, voyage[i] <= n`
- All the values in the tree are **unique**.
- All the values in `voyage` are **unique**.

## 题目大意

给你一棵二叉树的根节点 root ，树中有 n 个节点，每个节点都有一个不同于其他节点且处于 1 到 n 之间的值。另给你一个由 n 个值组成的行程序列 voyage ，表示 预期 的二叉树 先序遍历 结果。通过交换节点的左右子树，可以 翻转 该二叉树中的任意节点。请翻转 最少 的树中节点，使二叉树的 先序遍历 与预期的遍历行程 voyage 相匹配 。如果可以，则返回 翻转的 所有节点的值的列表。你可以按任何顺序返回答案。如果不能，则返回列表 [-1]。

## 解题思路

- 题目要求翻转最少树中节点，利用贪心的思想，应该从根节点开始从上往下依次翻转，这样翻转的次数是最少的。对树进行深度优先遍历，如果遍历到某一个节点的时候，节点值不能与行程序列匹配，那么答案一定是 [-1]。否则，当下一个期望数字 `voyage[i]` 与即将遍历的子节点的值不同的时候，就要翻转一下当前这个节点的左右子树，继续 DFS。递归结束可能有 2 种情况，一种是找出了所有要翻转的节点，另一种情况是没有需要翻转的，即原树先序遍历的结果与 `voyage` 是完全一致的。

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

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	res, index := make([]int, 0, len(voyage)), 0
	if travelTree(root, &index, voyage, &res) {
		return res
	}
	return []int{-1}
}

func travelTree(root *TreeNode, index *int, voyage []int, res *[]int) bool {
	if root == nil {
		return true
	}
	if root.Val != voyage[*index] {
		return false
	}
	*index++
	if root.Left != nil && root.Left.Val != voyage[*index] {
		*res = append(*res, root.Val)
		return travelTree(root.Right, index, voyage, res) && travelTree(root.Left, index, voyage, res)
	}
	return travelTree(root.Left, index, voyage, res) && travelTree(root.Right, index, voyage, res)
}
```