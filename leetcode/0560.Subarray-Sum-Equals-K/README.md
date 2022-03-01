# [560. Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/)


## 题目

Given an array of integers `nums` and an integer `k`, return *the total number of continuous subarrays whose sum equals to `k`*.

**Example 1:**

```
Input: nums = [1,1,1], k = 2
Output: 2

```

**Example 2:**

```
Input: nums = [1,2,3], k = 3
Output: 2

```

**Constraints:**

- `1 <= nums.length <= 2 * 104`
- `-1000 <= nums[i] <= 1000`
- `-10^7 <= k <= 10^7`

## 题目大意

给你一个整数数组 `nums` 和一个整数 `k` ，请你统计并返回该数组中和为 `k` ****的连续子数组的个数。

## 解题思路

- 此题不能使用滑动窗口来解。因为 `nums[i]` 可能为负数。
- 前缀和的思路可以解答此题，但是时间复杂度有点高了，`O(n^2)`。考虑优化时间复杂度。
- 题目要求找到连续区间和为 `k` 的子区间总数，即区间 `[i,j]` 内的和为 K ⇒ `prefixSum[j] - prefixSum[i-1] == k`。所以 `prefixSum[j] == k - prefixSum[i-1]` 。这样转换以后，题目就转换成类似 A + B = K 的问题了。LeetCode 第一题的优化思路拿来用。用 map 存储累加过的结果。如此优化以后，时间复杂度 `O(n)`。

## 代码

```go
package leetcode

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}
```