# [2164. Sort Even and Odd Indices Independently](https://leetcode.com/problems/sort-even-and-odd-indices-independently/)


## 题目

You are given a **0-indexed** integer array `nums`. Rearrange the values of `nums` according to the following rules:

1. Sort the values at **odd indices** of `nums` in **non-increasing** order.
    - For example, if `nums = [4,**1**,2,**3**]` before this step, it becomes `[4,**3**,2,**1**]` after. The values at odd indices `1` and `3` are sorted in non-increasing order.
2. Sort the values at **even indices** of `nums` in **non-decreasing** order.
    - For example, if `nums = [**4**,1,**2**,3]` before this step, it becomes `[**2**,1,**4**,3]` after. The values at even indices `0` and `2` are sorted in non-decreasing order.

Return *the array formed after rearranging the values of* `nums`.

**Example 1:**

```
Input: nums = [4,1,2,3]
Output: [2,3,4,1]
Explanation:
First, we sort the values present at odd indices (1 and 3) in non-increasing order.
So, nums changes from [4,1,2,3] to [4,3,2,1].
Next, we sort the values present at even indices (0 and 2) in non-decreasing order.
So, nums changes from [4,1,2,3] to [2,3,4,1].
Thus, the array formed after rearranging the values is [2,3,4,1].

```

**Example 2:**

```
Input: nums = [2,1]
Output: [2,1]
Explanation:
Since there is exactly one odd index and one even index, no rearrangement of values takes place.
The resultant array formed is [2,1], which is the same as the initial array.

```

**Constraints:**

- `1 <= nums.length <= 100`
- `1 <= nums[i] <= 100`

## 题目大意

给你一个下标从 0 开始的整数数组 nums 。根据下述规则重排 nums 中的值：

1. 按 非递增 顺序排列 nums 奇数下标 上的所有值。
举个例子，如果排序前 nums = [4,1,2,3] ，对奇数下标的值排序后变为 [4,3,2,1] 。奇数下标 1 和 3 的值按照非递增顺序重排。
2. 按 非递减 顺序排列 nums 偶数下标 上的所有值。
举个例子，如果排序前 nums = [4,1,2,3] ，对偶数下标的值排序后变为 [2,1,4,3] 。偶数下标 0 和 2 的值按照非递减顺序重排。

返回重排 nums 的值之后形成的数组。

## 解题思路

- 简单题。分别将奇数和偶数位上的数字排序，奇数位的数从大到小，偶数位的数从小到大。最后将他们组合成一个数组。

## 代码

```go
package leetcode

import (
	"sort"
)

func sortEvenOdd(nums []int) []int {
	odd, even, res := []int{}, []int{}, []int{}
	for index, v := range nums {
		if index%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	sort.Ints(even)
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))

	indexO, indexE := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			res = append(res, even[indexE])
			indexE++
		} else {
			res = append(res, odd[indexO])
			indexO++
		}
	}
	return res
}
```