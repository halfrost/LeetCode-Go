# [69. Sqrt(x)](https://leetcode.com/problems/sqrtx/)


## 题目

Implement `int sqrt(int x)`.

Compute and return the square root of *x*, where *x* is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

**Example 1:**

    Input: 4
    Output: 2

**Example 2:**

    Input: 8
    Output: 2
    Explanation: The square root of 8 is 2.82842..., and since 
                 the decimal part is truncated, 2 is returned.


## 题目大意

实现 int sqrt(int x) 函数。计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。



## 解题思路

- 题目要求求出根号 x
- 根据题意，根号 x 的取值范围一定在 `[0,x]` 之间，这个区间内的值是递增有序的，有边界的，可以用下标访问的，满足这三点正好也就满足了二分搜索的 3 大条件。所以解题思路一，二分搜索。
- 解题思路二，牛顿迭代法。求根号 x，即求满足 `x^2 - n = 0` 方程的所有解。

    ![](https://img-blog.csdn.net/20171019164040871?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvY2hlbnJlbnhpYW5n/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)
