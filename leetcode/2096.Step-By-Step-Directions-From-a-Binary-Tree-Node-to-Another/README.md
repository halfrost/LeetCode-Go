# [2096. Step-By-Step Directions From a Binary Tree Node to Another](https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/)


## 题目

You are given the `root` of a **binary tree** with `n` nodes. Each node is uniquely assigned a value from `1` to `n`. You are also given an integer `startValue` representing the value of the start node `s`, and a different integer `destValue` representing the value of the destination node `t`.

Find the **shortest path** starting from node `s` and ending at node `t`. Generate step-by-step directions of such path as a string consisting of only the **uppercase** letters `'L'`, `'R'`, and `'U'`. Each letter indicates a specific direction:

- `'L'` means to go from a node to its **left child** node.
- `'R'` means to go from a node to its **right child** node.
- `'U'` means to go from a node to its **parent** node.

Return *the step-by-step directions of the **shortest path** from node* `s` *to node* `t`.

**Example 1:**

![https://assets.leetcode.com/uploads/2021/11/15/eg1.png](https://assets.leetcode.com/uploads/2021/11/15/eg1.png)

```
Input: root = [5,1,2,3,null,6,4], startValue = 3, destValue = 6
Output: "UURL"
Explanation: The shortest path is: 3 → 1 → 5 → 2 → 6.

```

**Example 2:**

![https://assets.leetcode.com/uploads/2021/11/15/eg2.png](https://assets.leetcode.com/uploads/2021/11/15/eg2.png)

```
Input: root = [2,1], startValue = 2, destValue = 1
Output: "L"
Explanation: The shortest path is: 2 → 1.

```

**Constraints:**

- The number of nodes in the tree is `n`.
- `2 <= n <= 105`
- `1 <= Node.val <= n`
- All the values in the tree are **unique**.
- `1 <= startValue, destValue <= n`
- `startValue != destValue`

## 题目大意

给你一棵 二叉树 的根节点 root ，这棵二叉树总共有 n 个节点。每个节点的值为 1 到 n 中的一个整数，且互不相同。给你一个整数 startValue ，表示起点节点 s 的值，和另一个不同的整数 destValue ，表示终点节点 t 的值。

请找到从节点 s 到节点 t 的 最短路径 ，并以字符串的形式返回每一步的方向。每一步用 大写 字母 'L' ，'R' 和 'U' 分别表示一种方向：

- 'L' 表示从一个节点前往它的 左孩子 节点。
- 'R' 表示从一个节点前往它的 右孩子 节点。
- 'U' 表示从一个节点前往它的 父 节点。

请你返回从 s 到 t 最短路径 每一步的方向。

## 解题思路

- 二叉树中一个节点到另一个节点的最短路径一定可以分为两个部分（可能为空）：从起点节点向上到两个节点的**最近公共祖先**，再从最近公共祖先向下到达终点节点。
- 首先需要找到起点 s 与公共祖先的节点之间的 path1，公共祖先节点与终点 t 的 path2。再删掉 2 个 path 的公共前缀。如果起点 s 和终点 t 在不同的分支上，不存在公共前缀。如果他们在相同的分支上，那么最终答案要去掉这个公共前缀。
- 删除掉公共前缀以后，需要再整理一下最终答案的输出格式。由于题目要求，起点到公共祖先节点需要输出 U，所以把这段 path1 全部改成 U，然后再拼接上 path2 字符串，即可得到的字符串即为待求 ss 到 tt 每一步的最短路径。

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

func getDirections(root *TreeNode, startValue int, destValue int) string {
	sPath, dPath := make([]byte, 0), make([]byte, 0)
	findPath(root, startValue, &sPath)
	findPath(root, destValue, &dPath)
	size, i := min(len(sPath), len(dPath)), 0
	for i < size {
		if sPath[len(sPath)-1-i] == dPath[len(dPath)-1-i] {
			i++
		} else {
			break
		}
	}
	sPath = sPath[:len(sPath)-i]
	replace(sPath)
	dPath = dPath[:len(dPath)-i]
	reverse(dPath)
	sPath = append(sPath, dPath...)
	return string(sPath)
}

func findPath(root *TreeNode, value int, path *[]byte) bool {
	if root.Val == value {
		return true
	}

	if root.Left != nil && findPath(root.Left, value, path) {
		*path = append(*path, 'L')
		return true
	}

	if root.Right != nil && findPath(root.Right, value, path) {
		*path = append(*path, 'R')
		return true
	}

	return false
}

func reverse(path []byte) {
	left, right := 0, len(path)-1
	for left < right {
		path[left], path[right] = path[right], path[left]
		left++
		right--
	}
}

func replace(path []byte) {
	for i := 0; i < len(path); i++ {
		path[i] = 'U'
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
```