# [1305. All Elements in Two Binary Search Trees](https://leetcode.com/problems/all-elements-in-two-binary-search-trees/)



## 题目

Given two binary search trees `root1` and `root2`.

Return a list containing *all the integers* from *both trees* sorted in **ascending** order.

**Example 1**:

![https://assets.leetcode.com/uploads/2019/12/18/q2-e1.png](https://assets.leetcode.com/uploads/2019/12/18/q2-e1.png)

```
Input: root1 = [2,1,4], root2 = [1,0,3]
Output: [0,1,1,2,3,4]
```

**Example 2**:

```
Input: root1 = [0,-10,10], root2 = [5,1,7,0,2]
Output: [-10,0,0,1,2,5,7,10]
```

**Example 3**:

```
Input: root1 = [], root2 = [5,1,7,0,2]
Output: [0,1,2,5,7]
```

**Example 4**:

```
Input: root1 = [0,-10,10], root2 = []
Output: [-10,0,10]
```

**Example 5**:

![https://assets.leetcode.com/uploads/2019/12/18/q2-e5-.png](https://assets.leetcode.com/uploads/2019/12/18/q2-e5-.png)

```
Input: root1 = [1,null,8], root2 = [8,1]
Output: [1,1,8,8]
```

**Constraints**:

- Each tree has at most `5000` nodes.
- Each node's value is between `[-10^5, 10^5]`.

## 题目大意

给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.

提示：

- 每棵树最多有 5000 个节点。
- 每个节点的值在 [-10^5, 10^5] 之间。


## 解题思路

- 给出 2 棵二叉排序树，要求将 2 棵树所有节点的值按照升序排序。
- 这一题最暴力简单的方法就是把 2 棵树的节点都遍历出来，然后放在一个数组里面从小到大排序即可。这样做虽然能 AC，但是时间复杂度高。因为题目中给的二叉排序树这一条件没有用上。由于树中节点本来已经有序了，所以题目实质想要我们合并 2 个有序数组。利用中根遍历，把 2 个二叉排序树的所有节点值都遍历出来，遍历出来以后就是有序的。接下来再合并这两个有序数组即可。合并 2 个有序数组是第 88 题。

## 代码

```go
// 解法一 合并排序
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := inorderTraversal(root1)
	arr2 := inorderTraversal(root2)
	arr1 = append(arr1, make([]int, len(arr2))...)
	merge(arr1, len(arr1)-len(arr2), arr2, len(arr2))
	return arr1
}

// 解法二 暴力遍历排序，时间复杂度高
func getAllElements1(root1 *TreeNode, root2 *TreeNode) []int {
	arr := []int{}
	arr = append(arr, preorderTraversal(root1)...)
	arr = append(arr, preorderTraversal(root2)...)
	sort.Ints(arr)
	return arr
}
```