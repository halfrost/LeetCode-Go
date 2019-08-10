# [143. Reorder List](https://leetcode.com/problems/reorder-list/)

## 题目

Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:

```c
Given 1->2->3->4, reorder it to 1->4->2->3.
```

Example 2:

```c
Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
```

## 题目大意

按照指定规则重新排序链表：第一个元素和最后一个元素排列在一起，接着第二个元素和倒数第二个元素排在一起，接着第三个元素和倒数第三个元素排在一起。


## 解题思路


最近简单的方法是先把链表存储到数组里，然后找到链表中间的结点，按照规则拼接即可。这样时间复杂度是 O(n)，空间复杂度是 O(n)。

更好的做法是结合之前几道题的操作：链表逆序，找中间结点。

先找到链表的中间结点，然后利用逆序区间的操作，如 [第 92 题](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0092.%20Reverse%20Linked%20List%20II) 里的 reverseBetween() 操作，只不过这里的反转区间是从中点一直到末尾。最后利用 2 个指针，一个指向头结点，一个指向中间结点，开始拼接最终的结果。这种做法的时间复杂度是 O(n)，空间复杂度是 O(1)。