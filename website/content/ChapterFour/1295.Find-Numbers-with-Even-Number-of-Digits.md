# [1295. Find Numbers with Even Number of Digits](https://leetcode.com/problems/find-numbers-with-even-number-of-digits/)



## 题目

Given an array `nums` of integers, return how many of them contain an **even number** of digits.

**Example 1**:

```
Input: nums = [12,345,2,6,7896]
Output: 2
Explanation: 
12 contains 2 digits (even number of digits). 
345 contains 3 digits (odd number of digits). 
2 contains 1 digit (odd number of digits). 
6 contains 1 digit (odd number of digits). 
7896 contains 4 digits (even number of digits). 
Therefore only 12 and 7896 contain an even number of digits.
```

**Example 2**:

```
Input: nums = [555,901,482,1771]
Output: 1 
Explanation: 
Only 1771 contains an even number of digits.
```

**Constraints**:

- `1 <= nums.length <= 500`
- `1 <= nums[i] <= 10^5`

## 题目大意

给你一个整数数组 nums，请你返回其中位数为 偶数 的数字的个数。

提示：

- 1 <= nums.length <= 500
- 1 <= nums[i] <= 10^5



## 解题思路

- 给你一个整数数组，要求输出位数为偶数的数字的个数。
- 简单题，把每个数字转换为字符串判断长度是否是偶数即可。

## 代码

```go
func findNumbers(nums []int) int {
	res := 0
	for _, n := range nums {
		res += 1 - len(strconv.Itoa(n))%2
	}
	return res
}
```