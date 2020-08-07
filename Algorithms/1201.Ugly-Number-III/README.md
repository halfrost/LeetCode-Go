# [1201. Ugly Number III](https://leetcode.com/problems/ugly-number-iii/)


## 题目:

Write a program to find the `n`-th ugly number.

Ugly numbers are **positive integers** which are divisible by `a` **or** `b` **or** `c`.

**Example 1:**

    Input: n = 3, a = 2, b = 3, c = 5
    Output: 4
    Explanation: The ugly numbers are 2, 3, 4, 5, 6, 8, 9, 10... The 3rd is 4.

**Example 2:**

    Input: n = 4, a = 2, b = 3, c = 4
    Output: 6
    Explanation: The ugly numbers are 2, 3, 4, 6, 8, 9, 10, 12... The 4th is 6.

**Example 3:**

    Input: n = 5, a = 2, b = 11, c = 13
    Output: 10
    Explanation: The ugly numbers are 2, 4, 6, 8, 10, 11, 12, 13... The 5th is 10.

**Example 4:**

    Input: n = 1000000000, a = 2, b = 217983653, c = 336916467
    Output: 1999999984

**Constraints:**

- `1 <= n, a, b, c <= 10^9`
- `1 <= a * b * c <= 10^18`
- It's guaranteed that the result will be in range `[1, 2 * 10^9]`


## 题目大意


请你帮忙设计一个程序，用来找出第 n 个丑数。丑数是可以被 a 或 b 或 c 整除的 正整数。


提示：

- 1 <= n, a, b, c <= 10^9
- 1 <= a * b * c <= 10^18
- 本题结果在 [1, 2 * 10^9] 的范围内

## 解题思路


- 给出 4 个数字，a，b，c，n。要求输出可以整除 a 或者整除 b 或者整除 c 的第 n 个数。
- 这一题限定了解的范围， `[1, 2 * 10^9]`，所以直接二分搜索来求解。逐步二分逼近 low 值，直到找到能满足条件的 low 的最小值，即为最终答案。
- 这一题的关键在如何判断一个数是第几个数。一个数能整除 a，能整除 b，能整除 c，那么它应该是第 `num/a + num/b + num/c - num/ab - num/bc - num/ac + num/abc` 个数。这个就是韦恩图。需要注意的是，求 `ab`、`bc`、`ac`、`abc` 的时候需要再除以各自的最大公约数 `gcd()`。
