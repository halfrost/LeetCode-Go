# [540. Single Element in a Sorted Array](https://leetcode.com/problems/single-element-in-a-sorted-array/)

## 题目

You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.

Return the single element that appears only once.

Your solution must run in O(log n) time and O(1) space.

**Example 1:**

    Input: nums = [1,1,2,3,3,4,4,8,8]
    Output: 2

**Example 2:**

    Input: nums = [3,3,7,7,10,11,11]
    Output: 10

**Constraints:**

- 1 <= nums.length <= 100000
- 0 <= nums[i] <= 100000

## 题目大意

给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

## 解题思路

    假设下标idx是单独的数字,idx左边的偶数下标x有nums[x] == nums[x + 1],
    idx右边的奇数下标y有nums[y] == nums[y + 1],可以根据此特性用二分查找idx对应的值 

## 代码

```go
package leetcode

func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return nums[left]
}
```
