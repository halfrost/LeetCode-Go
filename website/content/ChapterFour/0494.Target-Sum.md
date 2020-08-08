# [494. Target Sum](https://leetcode.com/problems/target-sum/)


## 题目

You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols `+` and `-`. For each integer, you should choose one from `+` and `-` as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

**Example 1**:

```
Input: nums is [1, 1, 1, 1, 1], S is 3. 
Output: 5
Explanation: 

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
```

**Note**:

1. The length of the given array is positive and will not exceed 20.
2. The sum of elements in the given array will not exceed 1000.
3. Your output answer is guaranteed to be fitted in a 32-bit integer.

## 题目大意

给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在有两个符号 + 和 -。对于数组中的任意一个整数，可以从 + 或 - 中选择一个符号添加在前面。返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

提示：

- 数组非空，且长度不会超过 20 。
- 初始的数组的和不会超过 1000 。
- 保证返回的最终结果能被 32 位整数存下。

## 解题思路

- 给出一个数组，要求在这个数组里面的每个元素前面加上 + 或者 - 号，最终总和等于 S。问有多少种不同的方法。
- 这一题可以用 DP 和 DFS 解答。DFS 方法就不比较暴力简单了。见代码。这里分析一下 DP 的做法。题目要求在数组元素前加上 + 或者 - 号，其实相当于把数组分成了 2 组，一组全部都加 + 号，一组都加 - 号。记 + 号的一组 P ，记 - 号的一组 N，那么可以推出以下的关系。

    ```go
    sum(P) - sum(N) = target
    sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
                           2 * sum(P) = target + sum(nums)
    ```

    等号两边都加上 `sum(N) + sum(P)`，于是可以得到结果 `2 * sum(P) = target + sum(nums)`，那么这道题就转换成了，能否在数组中找到这样一个集合，和等于 `(target + sum(nums)) / 2`。那么这题就转化为了第 416 题了。`dp[i]` 中存储的是能使和为 `i` 的方法个数。

- 如果和不是偶数，即不能被 2 整除，那说明找不到满足题目要求的解了，直接输出 0 。

## 代码

```go

func findTargetSumWays(nums []int, S int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	if S > total || (S+total)%2 == 1 {
		return 0
	}
	target := (S + total) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, n := range nums {
		for i := target; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[target]
}

// 解法二 DFS
func findTargetSumWays1(nums []int, S int) int {
	// sums[i] 存储的是后缀和 nums[i:]，即从 i 到结尾的和
	sums := make([]int, len(nums))
	sums[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i > -1; i-- {
		sums[i] = sums[i+1] + nums[i]
	}
	res := 0
	dfsFindTargetSumWays(nums, 0, 0, S, &res, sums)
	return res
}

func dfsFindTargetSumWays(nums []int, index int, curSum int, S int, res *int, sums []int) {
	if index == len(nums) {
		if curSum == S {
			*(res) = *(res) + 1
		}
		return
	}
	// 剪枝优化：如果 sums[index] 值小于剩下需要正数的值，那么右边就算都是 + 号也无能为力了，所以这里可以剪枝了
	if S-curSum > sums[index] {
		return
	}
	dfsFindTargetSumWays(nums, index+1, curSum+nums[index], S, res, sums)
	dfsFindTargetSumWays(nums, index+1, curSum-nums[index], S, res, sums)
}

```