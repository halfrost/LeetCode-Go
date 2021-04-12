# [264. Ugly Number II](https://leetcode.com/problems/ugly-number-ii/)


## 题目

Given an integer `n`, return *the* `nth` ***ugly number***.

**Ugly number** is a positive number whose prime factors only include `2`, `3`, and/or `5`.

**Example 1:**

```
Input: n = 10
Output: 12
Explanation: [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] is the sequence of the first 10 ugly numbers.
```

**Example 2:**

```
Input: n = 1
Output: 1
Explanation: 1 is typically treated as an ugly number.
```

**Constraints:**

- `1 <= n <= 1690`

## 题目大意

给你一个整数 `n` ，请你找出并返回第 `n` 个 **丑数** 。**丑数** 就是只包含质因数 `2`、`3` 和/或 `5` 的正整数。

## 解题思路

- 解法一，生成丑数的方法：先用最小质因数 1，分别和 2，3，5 相乘，得到的数是丑数，不断的将这些数分别和 2，3，5 相乘，得到的数去重以后，从小到大排列，第 n 个数即为所求。排序可用最小堆实现，去重用 map 去重。时间复杂度 O(n log n)，空间复杂度 O(n)
- 上面的解法耗时在排序中，需要排序的根源是小的丑数乘以 5 大于了大的丑数乘以 2 。如何保证每次乘积以后，找出有序的丑数，是去掉排序，提升时间复杂度的关键。举个例子很容易想通：初始状态丑数只有 {1}，乘以 2，3，5 以后，将最小的结果存入集合中 {1,2}。下一轮再相乘，由于上一轮 1 已经和 2 相乘过了，1 不要再和 2 相乘了，所以这一轮 1 和 3，5 相乘。2 和 2，3，5 相乘。将最小的结果存入集合中 {1,2,3}，按照这样的策略往下比较，每轮选出的丑数是有序且不重复的。具体实现利用 3 个指针和一个数组即可实现。时间复杂度 O(n)，空间复杂度 O(n)。

## 代码

```go
package leetcode

func nthUglyNumber(n int) int {
	dp, p2, p3, p5 := make([]int, n+1), 1, 1, 1
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```