# [1696. Jump Game VI](https://leetcode.com/problems/jump-game-vi/)

## 题目

You are given a **0-indexed** integer array `nums` and an integer `k`.

You are initially standing at index `0`. In one move, you can jump at most `k` steps forward without going outside the boundaries of the array. That is, you can jump from index `i` to any index in the range `[i + 1, min(n - 1, i + k)]` **inclusive**.

You want to reach the last index of the array (index `n - 1`). Your **score** is the **sum** of all `nums[j]` for each index `j` you visited in the array.

Return *the **maximum score** you can get*.

**Example 1:**

```
Input: nums = [1,-1,-2,4,-7,3], k = 2
Output: 7
Explanation: You can choose your jumps forming the subsequence [1,-1,4,3] (underlined above). The sum is 7.

```

**Example 2:**

```
Input: nums = [10,-5,-2,4,0,3], k = 3
Output: 17
Explanation: You can choose your jumps forming the subsequence [10,4,3] (underlined above). The sum is 17.

```

**Example 3:**

```
Input: nums = [1,-5,-20,4,-1,3,-6,-3], k = 2
Output: 0

```

**Constraints:**

- `1 <= nums.length, k <= 10^5`
- `10^4 <= nums[i] <= 10^4`

## 题目大意

给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。一开始你在下标 0 处。每一步，你最多可以往前跳 k 步，但你不能跳出数组的边界。也就是说，你可以从下标 i 跳到 [i + 1， min(n - 1, i + k)] 包含 两个端点的任意位置。你的目标是到达数组最后一个位置（下标为 n - 1 ），你的 得分 为经过的所有数字之和。请你返回你能得到的 最大得分 。

## 解题思路

- 首先能想到的解题思路是动态规划。定义 `dp[i]` 为跳到第 `i` 个位子能获得的最大分数。题目要求的是 `dp[n-1]`，状态转移方程是：`dp[i] = nums[i] + max(dp[j]), max(0, i - k ) <= j < i`，这里需要注意 `j` 的下界，题目中说到不能跳到负数区间，所以左边界下界为 0 。求 `max(dp[j])` 需要遍历一次求得最大值，所以这个解法整体时间复杂度是 O((n - k) * k)，但是提交以后提示超时了。
- 分析一下超时原因。每次都要在 `[max(0, i - k ), i)` 区间内扫描找到最大值，下一轮的区间是 `[max(0, i - k + 1), i + 1)`，前后这两轮扫描的区间存在大量重合部分  `[max(0, i - k + 1), i)`，正是这部分反反复复的扫描导致算法低效。如何高效的在一个区间内找到最大值是本题的关键。利用单调队列可以完成此题。单调队列里面存一个区间内最大值的下标。这里单调队列有 2 个性质。性质一，队列的队首永远都是最大值，队列从大到小降序排列。如果来了一个比队首更大的值的下标，需要将单调队列清空，只存这个新的最大值的下标。性质二，队列的长度为 k。从队尾插入新值，并把队头的最大值“挤”出队首。拥有了这个单调队列以后，再进行 DP 状态转移，效率就很高了。每次只需取出队首的最大值即可。具体代码见下面。

## 代码

```go
package leetcode

import (
	"math"
)

// 单调队列
func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MinInt32
	}
	window := make([]int, k)
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i] + dp[window[0]]
		for len(window) > 0 && dp[window[len(window)-1]] <= dp[i] {
			window = window[:len(window)-1]
		}
		for len(window) > 0 && i-k >= window[0] {
			window = window[1:]
		}
		window = append(window, i)
	}
	return dp[len(nums)-1]
}

// 超时
func maxResult1(nums []int, k int) int {
	dp := make([]int, len(nums))
	if k > len(nums) {
		k = len(nums)
	}
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MinInt32
	}
	for i := 1; i < len(nums); i++ {
		left, tmp := max(0, i-k), math.MinInt32
		for j := left; j < i; j++ {
			tmp = max(tmp, dp[j])
		}
		dp[i] = nums[i] + tmp
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```