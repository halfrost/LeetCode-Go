# [518. Coin Change II](https://leetcode.com/problems/coin-change-ii/)


## 题目

You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return *the number of combinations that make up that amount*. If that amount of money cannot be made up by any combination of the coins, return `0`.

You may assume that you have an infinite number of each kind of coin.

The answer is **guaranteed** to fit into a signed **32-bit** integer.

**Example 1:**

```
Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
```

**Example 2:**

```
Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
```

**Example 3:**

```
Input: amount = 10, coins = [10]
Output: 1
```

**Constraints:**

- `1 <= coins.length <= 300`
- `1 <= coins[i] <= 5000`
- All the values of `coins` are **unique**.
- `0 <= amount <= 5000`

## 题目大意

给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。假设每一种面额的硬币有无限个。题目数据保证结果符合 32 位带符号整数。

## 解题思路

- 此题虽然名字叫 Coin Change，但是不是经典的背包九讲问题。题目中 coins 的每个元素可以选取多次，且不考虑选取元素的顺序，因此这道题实际需要计算的是选取硬币的组合数。定义 dp[i] 表示金额之和等于 i 的硬币组合数，目标求 dp[amount]。初始边界条件为 dp[0] = 1，即不取任何硬币，就这一种取法，金额为 0 。状态转移方程 dp[i] += dp[i-coin]，coin 为当前枚举的 coin。
- 可能有读者会有疑惑，上述做法会不会出现重复计算。答案是不会。外层循环是遍历数组 coins 的值，内层循环是遍历不同的金额之和，在计算 dp[i] 的值时，可以确保金额之和等于 i 的硬币面额的顺序，由于顺序确定，因此不会重复计算不同的排列。
- 和此题完全一致的解题思路的题有，第 377 题和第 494 题。

## 代码

```go
package leetcode

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
```