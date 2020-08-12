# [1464. Maximum Product of Two Elements in an Array](https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/)


## 题目

Given the array of integers `nums`, you will choose two different indices `i` and `j` of that array. Return the maximum value of `(nums[i]-1)*(nums[j]-1)`.

**Example 1**:

```
Input: nums = [3,4,5,2]
Output: 12 
Explanation: If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value, that is, (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12. 

```

**Example 2**:

```
Input: nums = [1,5,4,5]
Output: 16
Explanation: Choosing the indices i=1 and j=3 (indexed from 0), you will get the maximum value of (5-1)*(5-1) = 16.

```

**Example 3**:

```
Input: nums = [3,7]
Output: 12

```

**Constraints**:

- `2 <= nums.length <= 500`
- `1 <= nums[i] <= 10^3`

## 题目大意

给你一个整数数组 nums，请你选择数组的两个不同下标 i 和 j，使 (nums[i]-1)*(nums[j]-1) 取得最大值。请你计算并返回该式的最大值。

## 解题思路

- 简单题。循环一次，按照题意动态维护 2 个最大值，从而也使得 `(nums[i]-1)*(nums[j]-1)` 能取到最大值。

## 代码

```go

package leetcode

func maxProduct(nums []int) int {
	max1, max2 := 0, 0
	for _, num := range nums {
		if num >= max1 {
			max2 = max1
			max1 = num
		} else if num <= max1 && num >= max2 {
			max2 = num
		}
	}
	return (max1 - 1) * (max2 - 1)
}

```