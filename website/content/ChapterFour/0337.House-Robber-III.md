# [337. House Robber III](https://leetcode.com/problems/house-robber-iii/)



## 题目

The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

**Example 1**:

```
Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \ 
     3   1

Output: 7 
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
```

**Example 2**:

```
Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \ 
 1   3   1

Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
```

## 题目大意

一个新的可行窃的地区只有一个入口，称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。


## 解题思路

- 这一题是打家劫舍的第 3 题。这一题需要偷的房子是树状的。报警的条件还是相邻的房子如果都被偷了，就会触发报警。只不过这里相邻的房子是树上的。问小偷在不触发报警的条件下最终能偷的最高金额。
- 解题思路是 DFS。当前节点是否被打劫，会产生 2 种结果。如果当前节点被打劫，那么它的孩子节点可以被打劫；如果当前节点没有被打劫，那么它的孩子节点不能被打劫。按照这个逻辑递归，最终递归到根节点，取最大值输出即可。

## 代码

```go

func rob337(root *TreeNode) int {
	a, b := dfsTreeRob(root)
	return max(a, b)
}

func dfsTreeRob(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}
	l0, l1 := dfsTreeRob(root.Left)
	r0, r1 := dfsTreeRob(root.Right)
	// 当前节点没有被打劫
	tmp0 := max(l0, l1) + max(r0, r1)
	// 当前节点被打劫
	tmp1 := root.Val + l0 + r0
	return tmp0, tmp1
}

```