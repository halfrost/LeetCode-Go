# [2181. Merge Nodes in Between Zeros](https://leetcode.com/problems/merge-nodes-in-between-zeros/)

## 题目

You are given the `head` of a linked list, which contains a series of integers **separated** by `0`'s. The **beginning** and **end** of the linked list will have `Node.val == 0`.

For **every** two consecutive `0`'s, **merge** all the nodes lying in between them into a single node whose value is the **sum** of all the merged nodes. The modified list should not contain any `0`'s.

Return *the* `head` *of the modified linked list*.

**Example 1:**

![https://assets.leetcode.com/uploads/2022/02/02/ex1-1.png](https://assets.leetcode.com/uploads/2022/02/02/ex1-1.png)

```
Input: head = [0,3,1,0,4,5,2,0]
Output: [4,11]
Explanation:
The above figure represents the given linked list. The modified list contains
- The sum of the nodes marked in green: 3 + 1 = 4.
- The sum of the nodes marked in red: 4 + 5 + 2 = 11.

```

**Example 2:**

![https://assets.leetcode.com/uploads/2022/02/02/ex2-1.png](https://assets.leetcode.com/uploads/2022/02/02/ex2-1.png)

```
Input: head = [0,1,0,3,0,2,2,0]
Output: [1,3,4]
Explanation:
The above figure represents the given linked list. The modified list contains
- The sum of the nodes marked in green: 1 = 1.
- The sum of the nodes marked in red: 3 = 3.
- The sum of the nodes marked in yellow: 2 + 2 = 4.

```

**Constraints:**

- The number of nodes in the list is in the range `[3, 2 * 10^5]`.
- `0 <= Node.val <= 1000`
- There are **no** two consecutive nodes with `Node.val == 0`.
- The **beginning** and **end** of the linked list have `Node.val == 0`.

## 题目大意

给你一个链表的头节点 head ，该链表包含由 0 分隔开的一连串整数。链表的 开端 和 末尾 的节点都满足 Node.val == 0 。对于每两个相邻的 0 ，请你将它们之间的所有节点合并成一个节点，其值是所有已合并节点的值之和。然后将所有 0 移除，修改后的链表不应该含有任何 0 。

返回修改后链表的头节点 head 。

## 解题思路

- 简单题。合并链表中两个值为 0 的节点。从头开始遍历链表，遇到节点值不为 0 的节点便累加；遇到节点值为 0 的节点，将累加值转换成结果链表要输出的节点值，然后继续遍历。

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
func mergeNodes(head *ListNode) *ListNode {
	res := &ListNode{}
	h := res
	if head.Next == nil {
		return &structures.ListNode{}
	}
	cur := head
	sum := 0
	for cur.Next != nil {
		if cur.Next.Val != 0 {
			sum += cur.Next.Val
		} else {
			h.Next = &ListNode{Val: sum, Next: nil}
			h = h.Next
			sum = 0
		}
		cur = cur.Next
	}
	return res.Next
}
```