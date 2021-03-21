# [138. Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/)

## 题目

A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.

Return a **[deep copy](https://en.wikipedia.org/wiki/Object_copying#Deep_copy)** of the list.

The Linked List is represented in the input/output as a list of `n` nodes. Each node is represented as a pair of `[val, random_index]` where:

- `val`: an integer representing `Node.val`
- `random_index`: the index of the node (range from `0` to `n-1`) where random pointer points to, or `null` if it does not point to any node.

**Example 1:**

![](https://assets.leetcode.com/uploads/2019/12/18/e1.png)

```
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

```

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/12/18/e2.png)

```
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]

```

**Example 3:**

![](https://assets.leetcode.com/uploads/2019/12/18/e3.png)

```
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]

```

**Example 4:**

```
Input: head = []
Output: []
Explanation: Given linked list is empty (null pointer), so return null.

```

**Constraints:**

- `10000 <= Node.val <= 10000`
- `Node.random` is null or pointing to a node in the linked list.
- The number of nodes will not exceed 1000.

## 题目大意

给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。要求返回这个链表的 深拷贝。

我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

- val：一个表示 Node.val 的整数。
- random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。

## 解题思路

- 这道题严格意义上是数据结构题，根据给定的数据结构，对它进行深拷贝。
- 先将每个节点都复制一份，放在它的 next 节点中。如此穿插的复制一份链表。

    ![https://img.halfrost.com/Leetcode/leetcode_138_1_0.png](https://img.halfrost.com/Leetcode/leetcode_138_1_0.png)

    再将穿插版的链表的 random 指针指向正确的位置。

    ![https://img.halfrost.com/Leetcode/leetcode_138_2.png](https://img.halfrost.com/Leetcode/leetcode_138_2.png)

    再将穿插版的链表的 next 指针指向正确的位置。最后分开这交织在一起的两个链表的头节点，即可分开 2 个链表。

    ![https://img.halfrost.com/Leetcode/leetcode_138_3.png](https://img.halfrost.com/Leetcode/leetcode_138_3.png)

## 代码

```go
package leetcode

// Node define
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	tempHead := copyNodeToLinkedList(head)
	return splitLinkedList(tempHead)
}

func splitLinkedList(head *Node) *Node {
	cur := head
	head = head.Next
	for cur != nil && cur.Next != nil {
		cur.Next, cur = cur.Next.Next, cur.Next
	}
	return head
}

func copyNodeToLinkedList(head *Node) *Node {
	cur := head
	for cur != nil {
		node := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next, cur = node, cur.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	return head
}
```