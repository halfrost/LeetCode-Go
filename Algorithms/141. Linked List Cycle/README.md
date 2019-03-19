# [141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/description/)

## 题目

Given a linked list, determine if it has a cycle in it.

Follow up:
Can you solve it without using extra space?



## 题目大意

判断链表是否有环，不能使用额外的空间。

给 2 个指针，一个指针是另外一个指针的下一个指针。快指针一次走 2 格，慢指针一次走 1 格。如果存在环，那么前一个指针一定会经过若干圈之后追上慢的指针。