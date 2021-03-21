# [643. Maximum Average Subarray I](https://leetcode.com/problems/maximum-average-subarray-i/)

## 题目

Given an array consisting of `n` integers, find the contiguous subarray of given length `k` that has the maximum average value. And you need to output the maximum average value.

**Example 1:**

```
Input: [1,12,-5,-6,50,3], k = 4
Output: 12.75
Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
```

**Note:**

1. 1 <= `k` <= `n` <= 30,000.
2. Elements of the given array will be in the range [-10,000, 10,000].

## 题目大意

给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。

## 解题思路

- 简单题。循环一次，扫描数组过程中累加窗口大小为 k 的元素值。不断更新这个最大值。循环结束求出平均值即可。

## 代码

```go
package leetcode

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```