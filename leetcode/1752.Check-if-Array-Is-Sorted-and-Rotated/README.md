# [1752. Check if Array Is Sorted and Rotated](https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/)


## 题目

Given an array `nums`, return `true` *if the array was originally sorted in non-decreasing order, then rotated **some** number of positions (including zero)*. Otherwise, return `false`.

There may be **duplicates** in the original array.

**Note:** An array `A` rotated by `x` positions results in an array `B` of the same length such that `A[i] == B[(i+x) % A.length]`, where `%` is the modulo operation.

**Example 1:**

```
Input: nums = [3,4,5,1,2]
Output: true
Explanation: [1,2,3,4,5] is the original sorted array.
You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2].
```

**Example 2:**

```
Input: nums = [2,1,3,4]
Output: false
Explanation: There is no sorted array once rotated that can make nums.
```

**Example 3:**

```
Input: nums = [1,2,3]
Output: true
Explanation: [1,2,3] is the original sorted array.
You can rotate the array by x = 0 positions (i.e. no rotation) to make nums.
```

**Example 4:**

```
Input: nums = [1,1,1]
Output: true
Explanation: [1,1,1] is the original sorted array.
You can rotate any number of positions to make nums.
```

**Example 5:**

```
Input: nums = [2,1]
Output: true
Explanation: [1,2] is the original sorted array.
You can rotate the array by x = 5 positions to begin on the element of value 2: [2,1].
```

**Constraints:**

- `1 <= nums.length <= 100`
- `1 <= nums[i] <= 100`

## 题目大意

给你一个数组 nums 。nums 的源数组中，所有元素与 nums 相同，但按非递减顺序排列。如果 nums 能够由源数组轮转若干位置（包括 0 个位置）得到，则返回 true ；否则，返回 false 。源数组中可能存在 重复项 。

## 解题思路

- 简单题。从头扫描一遍数组，找出相邻两个元素递减的数对。如果递减的数对只有 1 个，则有可能是轮转得来的，超过 1 个，则返回 false。题干里面还提到可能有多个重复元素，针对这一情况还需要判断一下 `nums[0]` 和 `nums[len(nums)-1]` 。如果是相同元素，`nums[0] < nums[len(nums)-1]`，并且数组中间还存在一对递减的数对，这时候也是 false。判断好上述这 2 种情况，本题得解。

## 代码

```go
package leetcode

func check(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 || nums[0] < nums[len(nums)-1] {
				return false
			}
		}
	}
	return true
}
```