# [437. Path Sum III](https://leetcode.com/problems/path-sum-iii/)


## 题目

Given the `root` of a binary tree and an integer `targetSum`, return *the number of paths where the sum of the values along the path equals* `targetSum`.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

**Example 1:**

![https://assets.leetcode.com/uploads/2021/04/09/pathsum3-1-tree.jpg](https://assets.leetcode.com/uploads/2021/04/09/pathsum3-1-tree.jpg)

```
Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are shown.

```

**Example 2:**

```
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: 3

```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 1000]`.
- `109 <= Node.val <= 109`
- `1000 <= targetSum <= 1000`

## 题目大意

给定一个二叉树，它的每个结点都存放着一个整数值。找出路径和等于给定数值的路径总数。路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。


## 解题思路


- 这一题是第 112 题 Path Sum 和第 113 题 Path Sum II 的加强版，这一题要求求出任意一条路径的和为 sum，起点不一定是根节点，可以是任意节点开始。
- 注意这一题可能出现负数的情况，节点和为 sum，并不一定是最终情况，有可能下面还有正数节点和负数节点相加正好为 0，那么这也是一种情况。一定要遍历到底。
- 一个点是否为 sum 的起点，有 3 种情况，第一种情况路径包含该 root 节点，如果包含该结点，就在它的左子树和右子树中寻找和为 `sum-root.Val` 的情况。第二种情况路径不包含该 root 节点，那么就需要在它的左子树和右子树中分别寻找和为 sum 的结点。



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

// 解法一 带缓存 dfs
func pathSum(root *TreeNode, targetSum int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	return dfs(root, prefixSum, 0, targetSum)
}

func dfs(root *TreeNode, prefixSum map[int]int, cur, sum int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	cnt := 0
	if v, ok := prefixSum[cur-sum]; ok {
		cnt = v
	}
	prefixSum[cur]++
	cnt += dfs(root.Left, prefixSum, cur, sum)
	cnt += dfs(root.Right, prefixSum, cur, sum)
	prefixSum[cur]--
	return cnt
}

// 解法二
func pathSumIII(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := findPath437(root, sum)
	res += pathSumIII(root.Left, sum)
	res += pathSumIII(root.Right, sum)
	return res
}

// 寻找包含 root 这个结点，且和为 sum 的路径
func findPath437(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res++
	}
	res += findPath437(root.Left, sum-root.Val)
	res += findPath437(root.Right, sum-root.Val)
	return res
}

```