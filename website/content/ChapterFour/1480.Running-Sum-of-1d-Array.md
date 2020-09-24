# [1480. Running Sum of 1d Array](https://leetcode.com/problems/running-sum-of-1d-array/)

## 题目

Given an array `nums`. We define a running sum of an array as `runningSum[i] = sum(nums[0]…nums[i])`.

Return the running sum of `nums`.

**Example 1**:

```
Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

```

**Example 2**:

```
Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].

```

**Example 3**:

```
Input: nums = [3,1,2,10,1]
Output: [3,4,6,16,17]

```


**Constraints**:

- `1 <= nums.length <= 1000`
- `-10^6 <= nums[i] <= 10^6`

## 题目大意

给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。请返回 nums 的动态和。


## 解题思路

- 简单题，按照题意依次循环计算前缀和即可。

## 代码

```go
package leetcode

func runningSum(nums []int) []int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}
	return dp[1:]
}

```