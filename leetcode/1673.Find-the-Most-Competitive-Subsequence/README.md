# [1673. Find the Most Competitive Subsequence](https://leetcode.com/problems/find-the-most-competitive-subsequence/)


## 题目

Given an integer array `nums` and a positive integer `k`, return *the most **competitive** subsequence of* `nums` *of size* `k`.

An array's subsequence is a resulting sequence obtained by erasing some (possibly zero) elements from the array.

We define that a subsequence `a` is more **competitive** than a subsequence `b` (of the same length) if in the first position where `a` and `b` differ, subsequence `a` has a number **less** than the corresponding number in `b`. For example, `[1,3,4]` is more competitive than `[1,3,5]` because the first position they differ is at the final number, and `4` is less than `5`.

**Example 1:**

```
Input: nums = [3,5,2,6], k = 2
Output: [2,6]
Explanation: Among the set of every possible subsequence: {[3,5], [3,2], [3,6], [5,2], [5,6], [2,6]}, [2,6] is the most competitive.

```

**Example 2:**

```
Input: nums = [2,4,3,3,5,4,9,6], k = 4
Output: [2,3,3,4]

```

**Constraints:**

- `1 <= nums.length <= 105`
- `0 <= nums[i] <= 109`
- `1 <= k <= nums.length`

## 题目大意

给你一个整数数组 nums 和一个正整数 k ，返回长度为 k 且最具 竞争力 的 nums 子序列。数组的子序列是从数组中删除一些元素（可能不删除元素）得到的序列。

在子序列 a 和子序列 b 第一个不相同的位置上，如果 a 中的数字小于 b 中对应的数字，那么我们称子序列 a 比子序列 b（相同长度下）更具 竞争力 。 例如，[1,3,4] 比 [1,3,5] 更具竞争力，在第一个不相同的位置，也就是最后一个位置上， 4 小于 5 。

## 解题思路

- 这一题是单调栈的典型题型。利用单调栈，可以保证原数组中元素相对位置不变，这满足题意中删除元素但不移动元素的要求。单调栈又能保证每次进栈，元素是最小的。
- 类似的题目还有第 42 题，第 84 题，第 496 题，第 503 题，第 856 题，第 901 题，第 907 题，第 1130 题，第 1425 题，第 1673 题。

## 代码

```go
package leetcode

// 单调栈
func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack)+len(nums)-i > k && len(stack) > 0 && nums[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return stack[:k]
}
```