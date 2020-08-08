# [863. All Nodes Distance K in Binary Tree](https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/)



## 题目

We are given a binary tree (with root node `root`), a `target` node, and an integer value `K`.

Return a list of the values of all nodes that have a distance `K` from the `target` node. The answer can be returned in any order.

**Example 1**:

```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2

Output: [7,4,1]

Explanation: 
The nodes that are a distance 2 from the target node (with value 5)
have values 7, 4, and 1.
```

![https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png)

**Note**:

1. The given tree is non-empty.
2. Each node in the tree has unique values `0 <= node.val <= 500`.
3. The `target` node is a node in the tree.
4. `0 <= K <= 1000`.

## 题目大意

给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。

提示：

- 给定的树是非空的。
- 树上的每个结点都具有唯一的值 0 <= node.val <= 500 。
- 目标结点 target 是树上的结点。
- 0 <= K <= 1000.


## 解题思路

- 给出一颗树和一个目标节点 target，一个距离 K，要求找到所有距离目标节点 target 的距离是 K 的点。
- 这一题用 DFS 的方法解题。先找到当前节点距离目标节点的距离，如果在左子树中找到了 target，距离当前节点的距离 > 0，则还需要在它的右子树中查找剩下的距离。如果是在右子树中找到了 target，反之同理。如果当前节点就是目标节点，那么就可以直接记录这个点。否则每次遍历一个点，距离都减一。

## 代码

```go
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	visit := []int{}
	findDistanceK(root, target, K, &visit)
	return visit
}

func findDistanceK(root, target *TreeNode, K int, visit *[]int) int {
	if root == nil {
		return -1
	}
	if root == target {
		findChild(root, K, visit)
		return K - 1
	}
	leftDistance := findDistanceK(root.Left, target, K, visit)
	if leftDistance == 0 {
		findChild(root, leftDistance, visit)
	}
	if leftDistance > 0 {
		findChild(root.Right, leftDistance-1, visit)
		return leftDistance - 1
	}
	rightDistance := findDistanceK(root.Right, target, K, visit)
	if rightDistance == 0 {
		findChild(root, rightDistance, visit)
	}
	if rightDistance > 0 {
		findChild(root.Left, rightDistance-1, visit)
		return rightDistance - 1
	}
	return -1
}

func findChild(root *TreeNode, K int, visit *[]int) {
	if root == nil {
		return
	}
	if K == 0 {
		*visit = append(*visit, root.Val)
	} else {
		findChild(root.Left, K-1, visit)
		findChild(root.Right, K-1, visit)
	}
}
```