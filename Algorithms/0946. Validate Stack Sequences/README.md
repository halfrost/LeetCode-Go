# [946. Validate Stack Sequences](https://leetcode.com/problems/validate-stack-sequences/)

## 题目

Given two sequences pushed and popped with distinct values, return true if and only if this could have been the result of a sequence of push and pop operations on an initially empty stack.

 

Example 1:

```c
Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
Explanation: We might do the following sequence:
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

```

Example 2:

```c
Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
Output: false
Explanation: 1 cannot be popped before 2.
```

Note:

1. 0 <= pushed.length == popped.length <= 1000
2. 0 <= pushed[i], popped[i] < 1000
3. pushed is a permutation of popped.
4. pushed and popped have distinct values.

## 题目大意

给 2 个数组，一个数组里面代表的是 push 的顺序，另一个数组里面代表的是 pop 的顺序。问按照这样的顺序操作以后，最终能否把栈清空？

## 解题思路

这一题也是靠栈操作的题目，按照 push 数组的顺序先把压栈，然后再依次在 pop 里面找栈顶元素，找到了就 pop，直到遍历完 pop 数组，最终如果遍历完了 pop 数组，就代表清空了整个栈了。