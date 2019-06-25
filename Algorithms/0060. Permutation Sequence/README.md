# [60. Permutation Sequence](https://leetcode.com/problems/permutation-sequence/)


## 题目

The set `[1,2,3,...,*n*]` contains a total of *n*! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for *n* = 3:

1. `"123"`
2. `"132"`
3. `"213"`
4. `"231"`
5. `"312"`
6. `"321"`

Given *n* and *k*, return the *k*th permutation sequence.

**Note:**

- Given *n* will be between 1 and 9 inclusive.
- Given *k* will be between 1 and *n*! inclusive.

**Example 1:**

```c
Input: n = 3, k = 3
Output: "213"
```

**Example 2:**

```c
Input: n = 4, k = 9
Output: "2314"
```

## 题目大意

给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下："123"，"132"，"213"，"231"，"312"，"321"，给定 n 和 k，返回第 k 个排列。


## 解题思路

- 用 DFS 暴力枚举，这种做法时间复杂度特别高，想想更优的解法。
