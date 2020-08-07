# [509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)


## 题目:

The **Fibonacci numbers**, commonly denoted `F(n)` form a sequence, called the **Fibonacci sequence**, such that each number is the sum of the two preceding ones, starting from `0` and `1`. That is,

    F(0) = 0,   F(1) = 1
    F(N) = F(N - 1) + F(N - 2), for N > 1.

Given `N`, calculate `F(N)`.

**Example 1:**

    Input: 2
    Output: 1
    Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

**Example 2:**

    Input: 3
    Output: 2
    Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

**Example 3:**

    Input: 4
    Output: 3
    Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.

**Note:**

0 ≤ `N` ≤ 30.


## 题目大意

斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

```
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
```

给定 N，计算 F(N)。

提示：0 ≤ N ≤ 30

## 解题思路


- 求斐波那契数列
- 这一题解法很多，大的分类是四种，递归，记忆化搜索(dp)，矩阵快速幂，通项公式。其中记忆化搜索可以写 3 种方法，自底向上的，自顶向下的，优化空间复杂度版的。通项公式方法实质是求 a^b 这个还可以用快速幂优化时间复杂度到 O(log n) 。
