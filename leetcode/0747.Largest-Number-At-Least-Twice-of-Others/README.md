# [747. Largest Number At Least Twice of Others](https://leetcode.com/problems/largest-number-at-least-twice-of-others/)


## 题目

You are given an integer array `nums` where the largest integer is **unique**.

Determine whether the largest element in the array is **at least twice** as much as every other number in the array. If it is, return *the **index** of the largest element, or return* `-1` *otherwise*.

**Example 1:**

```
Input: nums = [3,6,1,0]
Output: 1
Explanation: 6 is the largest integer.
For every other number in the array x, 6 is at least twice as big as x.
The index of value 6 is 1, so we return 1.

```

**Example 2:**

```
Input: nums = [1,2,3,4]
Output: -1
Explanation: 4 is less than twice the value of 3, so we return -1.
```

**Example 3:**

```
Input: nums = [1]
Output: 0
Explanation: 1 is trivially at least twice the value as any other number because there are no other numbers.

```

**Constraints:**

- `1 <= nums.length <= 50`
- `0 <= nums[i] <= 100`
- The largest element in `nums` is unique.

## 题目大意

给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。

## 解题思路

- 简单题。先扫描一遍找到最大值和下标。再扫描一遍检查最大值是否是其他数字的两倍。

## 代码

```go
package leetcode

func dominantIndex(nums []int) int {
	maxNum, flag, index := 0, false, 0
	for i, v := range nums {
		if v > maxNum {
			maxNum = v
			index = i
		}
	}
	for _, v := range nums {
		if v != maxNum && 2*v > maxNum {
			flag = true
		}
	}
	if flag {
		return -1
	}
	return index
}
```