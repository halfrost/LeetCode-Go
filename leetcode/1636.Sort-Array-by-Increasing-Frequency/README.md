# [1636. Sort Array by Increasing Frequency](https://leetcode.com/problems/sort-array-by-increasing-frequency/)


## 题目

Given an array of integers `nums`, sort the array in **increasing** order based on the frequency of the values. If multiple values have the same frequency, sort them in **decreasing** order.

Return the *sorted array*.

**Example 1:**

```
Input: nums = [1,1,2,2,2,3]
Output: [3,1,1,2,2,2]
Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.
```

**Example 2:**

```
Input: nums = [2,3,1,3,2]
Output: [1,3,3,2,2]
Explanation: '2' and '3' both have a frequency of 2, so they are sorted in decreasing order.
```

**Example 3:**

```
Input: nums = [-1,1,-6,4,5,-6,1,4,1]
Output: [5,-1,4,4,-6,-6,1,1,1]
```

**Constraints:**

- `1 <= nums.length <= 100`
- `100 <= nums[i] <= 100`

## 题目大意

给你一个整数数组 `nums` ，请你将数组按照每个值的频率 **升序** 排序。如果有多个值的频率相同，请你按照数值本身将它们 **降序** 排序。请你返回排序后的数组。

## 解题思路

- 简单题。先统计每个值的频率，然后按照频率从小到大排序，相同频率的按照值的大小，从大到小排序。

## 代码

```go
package leetcode

import "sort"

func frequencySort(nums []int) []int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[j] < nums[i]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}
```