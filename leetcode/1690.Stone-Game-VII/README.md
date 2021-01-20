# [1690. Stone Game VII](https://leetcode.com/problems/stone-game-vii/)

## 题目

Alice and Bob take turns playing a game, with **Alice starting first**.

There are `n` stones arranged in a row. On each player's turn, they can **remove** either the leftmost stone or the rightmost stone from the row and receive points equal to the **sum** of the remaining stones' values in the row. The winner is the one with the higher score when there are no stones left to remove.

Bob found that he will always lose this game (poor Bob, he always loses), so he decided to **minimize the score's difference**. Alice's goal is to **maximize the difference** in the score.

Given an array of integers `stones` where `stones[i]` represents the value of the `ith` stone **from the left**, return *the **difference** in Alice and Bob's score if they both play **optimally**.*

**Example 1:**

```
Input: stones = [5,3,1,4,2]
Output: 6
Explanation: 
- Alice removes 2 and gets 5 + 3 + 1 + 4 = 13 points. Alice = 13, Bob = 0, stones = [5,3,1,4].
- Bob removes 5 and gets 3 + 1 + 4 = 8 points. Alice = 13, Bob = 8, stones = [3,1,4].
- Alice removes 3 and gets 1 + 4 = 5 points. Alice = 18, Bob = 8, stones = [1,4].
- Bob removes 1 and gets 4 points. Alice = 18, Bob = 12, stones = [4].
- Alice removes 4 and gets 0 points. Alice = 18, Bob = 12, stones = [].
The score difference is 18 - 12 = 6.
```

**Example 2:**

```
Input: stones = [7,90,5,1,100,10,10,2]
Output: 122
```

**Constraints:**

- `n == stones.length`
- `2 <= n <= 1000`
- `1 <= stones[i] <= 1000`

## 题目大意

石子游戏中，爱丽丝和鲍勃轮流进行自己的回合，爱丽丝先开始 。有 n 块石子排成一排。每个玩家的回合中，可以从行中 移除 最左边的石头或最右边的石头，并获得与该行中剩余石头值之 和 相等的得分。当没有石头可移除时，得分较高者获胜。鲍勃发现他总是输掉游戏（可怜的鲍勃，他总是输），所以他决定尽力 减小得分的差值 。爱丽丝的目标是最大限度地 扩大得分的差值 。

给你一个整数数组 stones ，其中 stones[i] 表示 从左边开始 的第 i 个石头的值，如果爱丽丝和鲍勃都 发挥出最佳水平 ，请返回他们 得分的差值 。

## 解题思路

- 首先考虑 Bob 缩小分值差距意味着什么，意味着他想让他和 Alice 相对分数最小。Bob 已经明确肯定是输，所以他的分数一定比 Alice 小，那么 Bob - Alice 分数相减一定是负数。相对分数越小，意味着差值越大。负数越大，差值越小。-50 和 -10，-10 数值大，相差小。所以 Bob 的操作是让相对分数越大。Alice 的目的也是这样，要让 Alice - Bob 的相对分数越大，这里是正数的越大。综上，两者的目的相同，都是让相对分数最大化。
- 定义 `dp[i][j]` 代表在当前 `stone[i ~ j]` 区间内能获得的最大分差。状态转移方程为：

    ```go
    dp[i][j] = max(
                    sum(i + 1, j) - dp[i + 1][j],   // 这一局取走 `stone[i]`，获得 sum(i + 1, j) 分数，再减去剩下对手能获得的分数，即是此局能获得的最大分差。
                    sum(i, j - 1) - dp[i][j - 1]    // 这一局取走 `stone[j]`，获得 sum(i, j - 1) 分数，再减去剩下对手能获得的分数，即是此局能获得的最大分差。
                  )
    ```

    计算 `sum(i + 1, j) = stone[i + 1] + stone[i + 2] + …… + stone[j]` 利用前缀和计算区间和。

- 解法二是正常思路解答出来的代码。解法一是压缩了 DP 数组，在 DP 状态转移的时候，生成下一个 `dp[j]` 实际上是有规律的。利用这个规律可以少存一维数据，压缩空间。解法一的代码直接写出来，比较难想。先写出解法二的代码，然后找到递推规律，优化空间压缩一维，再写出解法一的代码。

## 代码

```go
package leetcode

// 解法一 优化空间版 DP
func stoneGameVII(stones []int) int {
	n := len(stones)
	sum := make([]int, n)
	dp := make([]int, n)
	for i, d := range stones {
		sum[i] = d
	}
	for i := 1; i < n; i++ {
		for j := 0; j+i < n; j++ {
			if (n-i)%2 == 1 {
				d0 := dp[j] + sum[j]
				d1 := dp[j+1] + sum[j+1]
				if d0 > d1 {
					dp[j] = d0
				} else {
					dp[j] = d1
				}
			} else {
				d0 := dp[j] - sum[j]
				d1 := dp[j+1] - sum[j+1]
				if d0 < d1 {
					dp[j] = d0
				} else {
					dp[j] = d1
				}
			}
			sum[j] = sum[j] + stones[i+j]
		}
	}
	return dp[0]
}

// 解法二 常规 DP
func stoneGameVII1(stones []int) int {
	prefixSum := make([]int, len(stones))
	for i := 0; i < len(stones); i++ {
		if i == 0 {
			prefixSum[i] = stones[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + stones[i]
		}
	}
	dp := make([][]int, len(stones))
	for i := range dp {
		dp[i] = make([]int, len(stones))
		dp[i][i] = 0
	}
	n := len(stones)
	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			dp[i][i+l-1] = max(prefixSum[i+l-1]-prefixSum[i+1]+stones[i+1]-dp[i+1][i+l-1], prefixSum[i+l-2]-prefixSum[i]+stones[i]-dp[i][i+l-2])
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```