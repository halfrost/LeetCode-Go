# [665. Non-decreasing Array](https://leetcode.com/problems/non-decreasing-array/)

## 题目

Given an array `nums` with `n` integers, your task is to check if it could become non-decreasing by modifying **at most one element**.

We define an array is non-decreasing if `nums[i] <= nums[i + 1]` holds for every `i` (**0-based**) such that (`0 <= i <= n - 2`).

**Example 1:**

```
Input: nums = [4,2,3]
Output: true
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
```

**Example 2:**

```
Input: nums = [4,2,1]
Output: false
Explanation: You can't get a non-decreasing array by modify at most one element.
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 104`
- `-10^5 <= nums[i] <= 10^5`

## 题目大意

给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。我们是这样定义一个非递减数列的： 对于数组中任意的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

## 解题思路

- 简单题。循环扫描数组，找到 `nums[i] > nums[i+1]` 这种递减组合。一旦这种组合超过 2 组，直接返回 false。找到第一组递减组合，需要手动调节一次。如果 `nums[i + 1] < nums[i - 1]`，就算交换 `nums[i+1]` 和 `nums[i]`，交换结束，`nums[i - 1]` 仍然可能大于 `nums[i + 1]`，不满足题意。正确的做法应该是让较小的那个数变大，即 `nums[i + 1] = nums[i]`。两个元素相等满足非递减的要求。

## 代码

```go
package leetcode

func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 {
				return false
			}
			if i > 0 && nums[i+1] < nums[i-1] {
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}
```