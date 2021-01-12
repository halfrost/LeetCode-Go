# [1695. Maximum Erasure Value](https://leetcode.com/problems/maximum-erasure-value/)


## 题目

You are given an array of positive integers `nums` and want to erase a subarray containing **unique elements**. The **score** you get by erasing the subarray is equal to the **sum** of its elements.

Return *the **maximum score** you can get by erasing **exactly one** subarray.*

An array `b` is called to be a subarray of `a` if it forms a contiguous subsequence of `a`, that is, if it is equal to `a[l],a[l+1],...,a[r]` for some `(l,r)`.

**Example 1:**

```
Input: nums = [4,2,4,5,6]
Output: 17
Explanation: The optimal subarray here is [2,4,5,6].
```

**Example 2:**

```
Input: nums = [5,2,1,2,5,2,1,2,5]
Output: 8
Explanation: The optimal subarray here is [5,2,1] or [1,2,5].
```

**Constraints:**

- `1 <= nums.length <= 105`
- `1 <= nums[i] <= 104`

## 题目大意

给你一个正整数数组 nums ，请你从中删除一个含有 若干不同元素 的子数组。删除子数组的 得分 就是子数组各元素之 和 。返回 只删除一个 子数组可获得的 最大得分 。如果数组 b 是数组 a 的一个连续子序列，即如果它等于 a[l],a[l+1],...,a[r] ，那么它就是 a 的一个子数组。

## 解题思路

- 读完题立马能识别出这是经典的滑动窗口题。利用滑动窗口从左往右滑动窗口，滑动过程中统计频次，如果是不同元素，右边界窗口又移，否则左边窗口缩小。每次移动更新 max 值。最终扫完一遍以后，max 值即为所求。

## 代码

```go
package leetcode

func maximumUniqueSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result, left, right, freq := 0, 0, -1, map[int]int{}
	for left < len(nums) {
		if right+1 < len(nums) && freq[nums[right+1]] == 0 {
			freq[nums[right+1]]++
			right++
		} else {
			freq[nums[left]]--
			left++
		}
		sum := 0
		for i := left; i <= right; i++ {
			sum += nums[i]
		}
		result = max(result, sum)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
```