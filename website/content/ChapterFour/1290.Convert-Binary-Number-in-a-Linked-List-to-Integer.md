# [1290. Convert Binary Number in a Linked List to Integer](https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/)



## 题目

Given `head` which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.

Return the *decimal value* of the number in the linked list.

**Example 1**:

![https://assets.leetcode.com/uploads/2019/12/05/graph-1.png](https://assets.leetcode.com/uploads/2019/12/05/graph-1.png)

```
Input: head = [1,0,1]
Output: 5
Explanation: (101) in base 2 = (5) in base 10
```

**Example 2**:

```
Input: head = [0]
Output: 0
```

**Example 3**:

```
Input: head = [1]
Output: 1
```

**Example 4**:

```
Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
Output: 18880
```

**Example 5**:

```
Input: head = [0,0]
Output: 0
```

**Constraints**:

- The Linked List is not empty.
- Number of nodes will not exceed `30`.
- Each node's value is either `0` or `1`.

## 题目大意

给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。请你返回该链表所表示数字的 十进制值 。

提示：

- 链表不为空。
- 链表的结点总数不超过 30。
- 每个结点的值不是 0 就是 1。

## 解题思路

- 给出一个链表，链表从头到尾表示的数是一个整数的二进制形式，要求输出这个整数的十进制。
- 简单题，从头到尾遍历一次链表，边遍历边累加二进制位。

## 代码

```go
func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum*2 + head.Val
		head = head.Next
	}
	return sum
}
```