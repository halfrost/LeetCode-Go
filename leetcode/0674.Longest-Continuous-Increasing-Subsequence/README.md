# [674. Longest Continuous Increasing Subsequence](https://leetcode.com/problems/longest-continuous-increasing-subsequence/)


## 题目

Given an unsorted array of integers `nums`, return *the length of the longest **continuous increasing subsequence** (i.e. subarray)*. The subsequence must be **strictly** increasing.

A **continuous increasing subsequence** is defined by two indices `l` and `r` (`l < r`) such that it is `[nums[l], nums[l + 1], ..., nums[r - 1], nums[r]]` and for each `l <= i < r`, `nums[i] < nums[i + 1]`.

**Example 1:**

```
Input: nums = [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
Even though [1,3,5,7] is an increasing subsequence, it is not continuous as elements 5 and 7 are separated by element
4.
```

**Example 2:**

```
Input: nums = [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2] with length 1. Note that it must be strictly
increasing.
```

**Constraints:**

- `0 <= nums.length <= 10^4`
- `10^9 <= nums[i] <= 10^9`

## 题目大意

给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

## 解题思路

- 简单题。这一题和第 128 题有区别。这一题要求子序列必须是连续下标，所以变简单了。扫描一遍数组，记下连续递增序列的长度，动态维护这个最大值，最后输出即可。

## 代码

```go
package leetcode

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, length := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			length++
		} else {
			res = max(res, length)
			length = 1
		}
	}
	return max(res, length)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```