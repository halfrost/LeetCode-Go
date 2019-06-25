# [326. Power of Three](https://leetcode.com/problems/power-of-three/)


## 题目

Given an integer, write a function to determine if it is a power of three.

**Example 1:**

    Input: 27
    Output: true

**Example 2:**

    Input: 0
    Output: false

**Example 3:**

    Input: 9
    Output: true

**Example 4:**

    Input: 45
    Output: false

**Follow up:**

Could you do it without using any loop / recursion?


## 题目大意

给定一个整数，写一个函数来判断它是否是 3 的幂次方。


## 解题思路

- 判断一个数是不是 3 的 n 次方。
- 这一题最简单的思路是循环，可以通过。但是题目要求不循环就要判断，这就需要用到数论的知识了。由于 3^20 超过了 int 的范围了，所以 3^19 次方就是 int 类型中最大的值。这一题和第 231 题是一样的思路。

