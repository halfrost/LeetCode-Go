# [234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)

## 题目

Given a singly linked list, determine if it is a palindrome.

Example 1:

```c
Input: 1->2
Output: false
```

Example 2:

```c
Input: 1->2->2->1
Output: true
```

Follow up:  

Could you do it in O(n) time and O(1) space?

## 题目大意

判断一个链表是否是回文链表。要求时间复杂度 O(n)，空间复杂度 O(1)。

## 解题思路

这道题只需要在第 143 题上面改改就可以了。思路是完全一致的。先找到中间结点，然后反转中间结点后面到结尾的所有结点。最后一一判断头结点开始的结点和中间结点往后开始的结点是否相等。如果一直相等，就是回文链表，如果有不相等的，直接返回不是回文链表。