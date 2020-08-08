# [1302. Deepest Leaves Sum](https://leetcode.com/problems/deepest-leaves-sum/)



## 题目

Given a binary tree, return the sum of values of its deepest leaves.

**Example 1**:

![https://assets.leetcode.com/uploads/2019/07/31/1483_ex1.png](https://assets.leetcode.com/uploads/2019/07/31/1483_ex1.png)

```
Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
Output: 15
```

**Constraints**:

- The number of nodes in the tree is between `1` and `10^4`.
- The value of nodes is between `1` and `100`.

## 题目大意

给你一棵二叉树，请你返回层数最深的叶子节点的和。

提示：

- 树中节点数目在 1 到 10^4 之间。
- 每个节点的值在 1 到 100 之间。

## 解题思路

- 给你一棵二叉树，请你返回层数最深的叶子节点的和。
- 这一题不难，DFS 遍历把最底层的叶子节点和都加起来即可。

## 代码

```go
func deepestLeavesSum(root *TreeNode) int {
	maxLevel, sum := 0, 0
	dfsDeepestLeavesSum(root, 0, &maxLevel, &sum)
	return sum
}

func dfsDeepestLeavesSum(root *TreeNode, level int, maxLevel, sum *int) {
	if root == nil {
		return
	}
	if level > *maxLevel {
		*maxLevel, *sum = level, root.Val
	} else if level == *maxLevel {
		*sum += root.Val
	}
	dfsDeepestLeavesSum(root.Left, level+1, maxLevel, sum)
	dfsDeepestLeavesSum(root.Right, level+1, maxLevel, sum)
}
```