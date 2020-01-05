# [1171. Remove Zero Sum Consecutive Nodes from Linked List](https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/)


## 题目:

Given the `head` of a linked list, we repeatedly delete consecutive sequences of nodes that sum to `0` until there are no such sequences.

After doing so, return the head of the final linked list. You may return any such answer.

(Note that in the examples below, all sequences are serializations of `ListNode` objects.)

**Example 1:**

    Input: head = [1,2,-3,3,1]
    Output: [3,1]
    Note: The answer [1,2,1] would also be accepted.

**Example 2:**

    Input: head = [1,2,3,-3,4]
    Output: [1,2,4]

**Example 3:**

    Input: head = [1,2,3,-3,-2]
    Output: [1]

**Constraints:**

- The given linked list will contain between `1` and `1000` nodes.
- Each node in the linked list has `-1000 <= node.val <= 1000`.


## 题目大意


给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。删除完毕后，请你返回最终结果链表的头节点。你可以返回任何满足题目要求的答案。

（注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）

提示：

- 给你的链表中可能有 1 到 1000 个节点。
- 对于链表中的每个节点，节点的值：-1000 <= node.val <= 1000.



## 解题思路

- 给出一个链表，要求把链表中和为 0 的结点都移除。
- 由于链表的特性，不能随机访问。所以从链表的头开始往后扫，把累加和存到字典中。当再次出现相同的累加和的时候，代表这中间的一段和是 0，于是要删除这一段。删除这一段的过程中，也要删除这一段在字典中存过的累加和。有一个特殊情况需要处理，即整个链表的总和是 0，那么最终结果是空链表。针对这个特殊情况，字典中先预存入 0 值。
