# [1104. Path In Zigzag Labelled Binary Tree](https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/)


## 题目

In an infinite binary tree where every node has two children, the nodes are labelled in row order.

In the odd numbered rows (ie., the first, third, fifth,...), the labelling is left to right, while in the even numbered rows (second, fourth, sixth,...), the labelling is right to left.

![https://assets.leetcode.com/uploads/2019/06/24/tree.png](https://assets.leetcode.com/uploads/2019/06/24/tree.png)

Given the `label` of a node in this tree, return the labels in the path from the root of the tree to the node with that `label`.

**Example 1:**

```
Input: label = 14
Output: [1,3,4,14]

```

**Example 2:**

```
Input: label = 26
Output: [1,2,6,10,26]

```

**Constraints:**

- `1 <= label <= 10^6`

## 题目大意

在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。

给你树上某一个节点的标号 label，请你返回从根节点到该标号为 label 节点的路径，该路径是由途经的节点标号所组成的。

## 解题思路

- 计算出 label 所在的 level 和 index。
- 根据 index 和 level 计算出父节点的 index 和 value。
- level 减一，循环计算出对应的父节点直到根节点。

## 代码

```go
package leetcode

func pathInZigZagTree(label int) []int {
	level := getLevel(label)
	ans := []int{label}
	curIndex := label - (1 << level)
	parent := 0
	for level >= 1 {
		parent, curIndex = getParent(curIndex, level)
		ans = append(ans, parent)
		level--
	}
	ans = reverse(ans)
	return ans
}

func getLevel(label int) int {
	level := 0
	nums := 0
	for {
		nums += 1 << level
		if nums >= label {
			return level
		}
		level++
	}
}

func getParent(index int, level int) (parent int, parentIndex int) {
	parentIndex = 1<<(level-1) - 1 + (index/2)*(-1)
	parent = 1<<(level-1) + parentIndex
	return
}

func reverse(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
```