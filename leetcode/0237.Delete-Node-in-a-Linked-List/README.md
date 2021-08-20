# [237. Delete Node in a Linked List](https://leetcode.com/problems/delete-node-in-a-linked-list/)


## 题目

Write a function to **delete a node** in a singly-linked list. You will **not** be given access to the `head` of the list, instead you will be given access to **the node to be deleted** directly.

It is **guaranteed** that the node to be deleted is **not a tail node** in the list.

**Example 1:**

![https://assets.leetcode.com/uploads/2020/09/01/node1.jpg](https://assets.leetcode.com/uploads/2020/09/01/node1.jpg)

```
Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation:You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/09/01/node2.jpg](https://assets.leetcode.com/uploads/2020/09/01/node2.jpg)

```
Input: head = [4,5,1,9], node = 1
Output: [4,5,9]
Explanation:You are given the third node with value 1, the linked list should become 4 -> 5 -> 9 after calling your function.

```

**Example 3:**

```
Input: head = [1,2,3,4], node = 3
Output: [1,2,4]

```

**Example 4:**

```
Input: head = [0,1], node = 0
Output: [1]

```

**Example 5:**

```
Input: head = [-3,5,-99], node = -3
Output: [5,-99]

```

**Constraints:**

- The number of the nodes in the given list is in the range `[2, 1000]`.
- `1000 <= Node.val <= 1000`
- The value of each node in the list is **unique**.
- The `node` to be deleted is **in the list** and is **not a tail** node

## 题目大意

删除给点结点。没有给链表的头结点。

## 解题思路

其实就是把后面的结点都覆盖上来即可。或者直接当前结点的值等于下一个结点，Next 指针指向下下个结点，这样做也可以，只不过中间有一个结点不被释放，内存消耗多一些。

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
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
```