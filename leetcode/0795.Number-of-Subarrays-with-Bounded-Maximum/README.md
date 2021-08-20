# [795. Number of Subarrays with Bounded Maximum](https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/)


## 题目

We are given an array `nums` of positive integers, and two positive integers `left` and `right` (`left <= right`).

Return the number of (contiguous, non-empty) subarrays such that the value of the maximum array element in that subarray is at least `left` and at most `right`.

```
Example:Input:
nums = [2, 1, 4, 3]
left = 2
right = 3
Output: 3
Explanation: There are three subarrays that meet the requirements: [2], [2, 1], [3].
```

**Note:**

- `left`, `right`, and `nums[i]` will be an integer in the range `[0, 109]`.
- The length of `nums` will be in the range of `[1, 50000]`.

## 题目大意

给定一个元素都是正整数的数组`A` ，正整数 `L` 以及 `R` (`L <= R`)。求连续、非空且其中最大元素满足大于等于`L` 小于等于`R`的子数组个数。

## 解题思路

- 题目要求子数组最大元素在 [L,R] 区间内。假设 count(bound) 为计算所有元素都小于等于 bound 的子数组数量。那么本题所求的答案可转化为 count(R) - count(L-1)。
- 如何统计所有元素小于 bound 的子数组数量呢？使用 count 变量记录在 bound 的左边，小于等于 bound 的连续元素数量。当找到一个这样的元素时，在此位置上结束的有效子数组的数量为 count + 1。当遇到一个元素大于 B 时，则在此位置结束的有效子数组的数量为 0。res 将每轮 count 累加，最终 res 中存的即是满足条件的所有子数组数量。

## 代码

```go
package leetcode

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	return getAnswerPerBound(nums, right) - getAnswerPerBound(nums, left-1)
}

func getAnswerPerBound(nums []int, bound int) int {
	res, count := 0, 0
	for _, num := range nums {
		if num <= bound {
			count++
		} else {
			count = 0
		}
		res += count
	}
	return res
}
```