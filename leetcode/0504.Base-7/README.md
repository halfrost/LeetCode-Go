# [504. Base 7](https://leetcode.com/problems/base-7/)

## 题目

Given an integer num, return a string of its base 7 representation.

**Example 1:**

    Input: num = 100
    Output: "202"

**Example 2:**

    Input: num = -7
    Output: "-10"

**Constraints:**

- -10000000 <= num <= 10000000

## 题目大意

给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。

## 解题思路

    num反复除以7，然后倒排余数

# 代码

```go
package leetcode

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negative := false
	if num < 0 {
		negative = true
		num = -num
	}
	var ans string
	var nums []int
	for num != 0 {
		remainder := num % 7
		nums = append(nums, remainder)
		num = num / 7
	}
	if negative {
		ans += "-"
	}
	for i := len(nums) - 1; i >= 0; i-- {
		ans += strconv.Itoa(nums[i])
	}
	return ans
}
```