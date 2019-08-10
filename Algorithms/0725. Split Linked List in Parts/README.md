# [725. Split Linked List in Parts](https://leetcode.com/problems/split-linked-list-in-parts/)

## 题目

Given a (singly) linked list with head node root, write a function to split the linked list into k consecutive linked list "parts".

The length of each part should be as equal as possible: no two parts should have a size differing by more than 1. This may lead to some parts being null.

The parts should be in order of occurrence in the input list, and parts occurring earlier should always have a size greater than or equal parts occurring later.

Return a List of ListNode's representing the linked list parts that are formed.

Examples 1->2->3->4, k = 5 // 5 equal parts [ [1], [2], [3], [4], null ]

Example 1:

```c
Input: 
root = [1, 2, 3], k = 5
Output: [[1],[2],[3],[],[]]
Explanation:
The input and each element of the output are ListNodes, not arrays.
For example, the input root has root.val = 1, root.next.val = 2, \root.next.next.val = 3, and root.next.next.next = null.
The first element output[0] has output[0].val = 1, output[0].next = null.
The last element output[4] is null, but it's string representation as a ListNode is [].
```

Example 2:

```c
Input: 
root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
Output: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
Explanation:
The input has been split into consecutive parts with size difference at most 1, and earlier parts are a larger size than the later parts.
```

Note:

- The length of root will be in the range [0, 1000].
- Each value of a node in the input will be an integer in the range [0, 999].
- k will be an integer in the range [1, 50].



## 题目大意

把链表分成 K 个部分，要求这 K 个部分尽量两两长度相差不超过 1，并且长度尽量相同。

## 解题思路

把链表长度对 K 进行除法，结果就是最终每组的长度 n。把链表长度对 K 进行取余操作，得到的结果 m，代表前 m 组链表长度为 n + 1 。相当于把多出来的部分都分摊到前面 m 组链表中了。最终链表是前 m 组长度为 n + 1，后 K - m 组链表长度是 n。

注意长度不足 K 的时候要用 nil 进行填充。



