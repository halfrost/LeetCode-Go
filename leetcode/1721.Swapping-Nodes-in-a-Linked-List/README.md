# [1721. Swapping Nodes in a Linked List](https://leetcode.com/problems/swapping-nodes-in-a-linked-list/)


## 题目

You are given the `head` of a linked list, and an integer `k`.

Return *the head of the linked list after **swapping** the values of the* `kth` *node from the beginning and the* `kth` *node from the end (the list is **1-indexed**).*

**Example 1:**

![https://assets.leetcode.com/uploads/2020/09/21/linked1.jpg](https://assets.leetcode.com/uploads/2020/09/21/linked1.jpg)

```
Input: head = [1,2,3,4,5], k = 2
Output: [1,4,3,2,5]
```

**Example 2:**

```
Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
Output: [7,9,6,6,8,7,3,0,9,5]
```

**Example 3:**

```
Input: head = [1], k = 1
Output: [1]
```

**Example 4:**

```
Input: head = [1,2], k = 1
Output: [2,1]
```

**Example 5:**

```
Input: head = [1,2,3], k = 2
Output: [1,2,3]
```

**Constraints:**

- The number of nodes in the list is `n`.
- `1 <= k <= n <= 10^5`
- `0 <= Node.val <= 100`

## 题目大意

给你链表的头节点 `head` 和一个整数 `k` 。**交换** 链表正数第 `k` 个节点和倒数第 `k` 个节点的值后，返回链表的头节点（链表 **从 1 开始索引**）。

## 解题思路

- 这道题虽然是 medium，但是实际非常简单。题目要求链表中 2 个节点的值，无非是先找到这 2 个节点，然后再交换即可。链表查询节点需要 O(n)，2 次循环找到对应的 2 个节点，交换值即可。

## 代码

```go
package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	count := 1
	var a, b *ListNode
	for node := head; node != nil; node = node.Next {
		if count == k {
			a = node
		}
		count++
	}
	length := count
	count = 1
	for node := head; node != nil; node = node.Next {
		if count == length-k {
			b = node
		}
		count++
	}
	a.Val, b.Val = b.Val, a.Val
	return head
}
```