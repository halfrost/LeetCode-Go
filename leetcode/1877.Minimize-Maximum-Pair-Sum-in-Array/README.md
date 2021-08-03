# [1877. Minimize Maximum Pair Sum in Array](https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/)


## 题目

The **pair sum** of a pair `(a,b)` is equal to `a + b`. The **maximum pair sum** is the largest **pair sum** in a list of pairs.

- For example, if we have pairs `(1,5)`, `(2,3)`, and `(4,4)`, the **maximum pair sum** would be `max(1+5, 2+3, 4+4) = max(6, 5, 8) = 8`.

Given an array `nums` of **even** length `n`, pair up the elements of `nums` into `n / 2` pairs such that:

- Each element of `nums` is in **exactly one** pair, and
- The **maximum pair sum** is **minimized**.

Return *the minimized **maximum pair sum** after optimally pairing up the elements*.

**Example 1:**

```
Input: nums = [3,5,2,3]
Output: 7
Explanation: The elements can be paired up into pairs (3,3) and (5,2).
The maximum pair sum is max(3+3, 5+2) = max(6, 7) = 7.
```

**Example 2:**

```
Input: nums = [3,5,4,2,4,6]
Output: 8
Explanation: The elements can be paired up into pairs (3,5), (4,4), and (6,2).
The maximum pair sum is max(3+5, 4+4, 6+2) = max(8, 8, 8) = 8.
```

**Constraints:**

- `n == nums.length`
- `2 <= n <= 105`
- `n` is **even**.
- `1 <= nums[i] <= 105`

## 题目大意

一个数对 (a,b) 的 **数对和** 等于 a + b 。**最大数对和** 是一个数对数组中最大的 数对和 。

- 比方说，如果我们有数对 (1,5) ，(2,3) 和 (4,4)，**最大数对和** 为 max(1+5, 2+3, 4+4) = max(6, 5, 8) = 8 。

给你一个长度为 **偶数** n 的数组 nums ，请你将 nums 中的元素分成 n / 2 个数对，使得：

- nums 中每个元素 **恰好** 在 一个 数对中，且
- **最大数对和** 的值 **最小** 。

请你在最优数对划分的方案下，返回最小的 最大数对和 。

## 解题思路

- 要想最大数对和最小，那么最大的元素一定只能和最小的元素组合在一起，不然一定不是最小。当最大元素和最小元素组合在一起了，剩下的次最大元素也应该和次最小元素组合在一起。按照这个思路，先将数组从小到大排序，然后依次取出首尾元素，两两组合在一起。输出这些数对的最大值即为所求。

## 代码

```go
package leetcode

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n, res := len(nums), 0
	for i, val := range nums[:n/2] {
		res = max(res, val+nums[n-1-i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```