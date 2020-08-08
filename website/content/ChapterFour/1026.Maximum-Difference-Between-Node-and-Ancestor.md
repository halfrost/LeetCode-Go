# [1026. Maximum Difference Between Node and Ancestor](https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/)



## 题目

Given the `root` of a binary tree, find the maximum value `V` for which there exists **different** nodes `A` and `B` where `V = |A.val - B.val|` and `A` is an ancestor of `B`.

(A node A is an ancestor of B if either: any child of A is equal to B, or any child of A is an ancestor of B.)

**Example 1**:

![https://assets.leetcode.com/uploads/2019/09/09/2whqcep.jpg](https://assets.leetcode.com/uploads/2019/09/09/2whqcep.jpg)

```
Input: [8,3,10,1,6,null,14,null,null,4,7,13]
Output: 7
Explanation: 
We have various ancestor-node differences, some of which are given below :
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.
```

**Note**:

1. The number of nodes in the tree is between `2` and `5000`.
2. Each node will have value between `0` and `100000`.

## 题目大意

给定二叉树的根节点 root，找出存在于不同节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）

提示：

- 树中的节点数在 2 到 5000 之间。
- 每个节点的值介于 0 到 100000 之间。



## 解题思路

- 给出一颗树，要求找出祖先和孩子的最大差值。
- DPS 深搜即可。每个节点和其所有孩子的`最大值`来自于 3 个值，节点本身，递归遍历左子树的最大值，递归遍历右子树的最大值；每个节点和其所有孩子的`最小值`来自于 3 个值，节点本身，递归遍历左子树的最小值，递归遍历右子树的最小值。依次求出自身节点和其所有孩子节点的最大差值，深搜的过程中动态维护最大差值即可。

## 代码

```go
func maxAncestorDiff(root *TreeNode) int {
	res := 0
	dfsAncestorDiff(root, &res)
	return res
}

func dfsAncestorDiff(root *TreeNode, res *int) (int, int) {
	if root == nil {
		return -1, -1
	}
	leftMax, leftMin := dfsAncestorDiff(root.Left, res)
	if leftMax == -1 && leftMin == -1 {
		leftMax = root.Val
		leftMin = root.Val
	}
	rightMax, rightMin := dfsAncestorDiff(root.Right, res)
	if rightMax == -1 && rightMin == -1 {
		rightMax = root.Val
		rightMin = root.Val
	}
	*res = max(*res, max(abs(root.Val-min(leftMin, rightMin)), abs(root.Val-max(leftMax, rightMax))))
	return max(leftMax, max(rightMax, root.Val)), min(leftMin, min(rightMin, root.Val))
}
```