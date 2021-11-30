# [400. Nth Digit](https://leetcode-cn.com/problems/nth-digit/)

## 题目

Given an integer n, return the nth digit of the infinite integer sequence [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...].

**Example 1**:

    Input: n = 3
    Output: 3

**Example 2**:

    Input: n = 11
    Output: 0
    Explanation: The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.

**Constraints:**

- 1 <= n <= int(math.Pow(2, 31)) - 1

## 题目大意

给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...] 中找出并返回第 n 位数字。

## 解题思路

- bits = 1的时候有1,2,3,4,5,6,7,8,9这9个数;9 = math.Pow10(bits - 1) * bits
- bits = 2的时候有10-99这90个数;90 = math.Pow10(bits - 1) * bits
- n不断减去bits从1开始的数字总数,求出n所在的数字是几位数即bits
- 计算n所在的数字num，等于初始值加上(n - 1) / bits
- 计算n所在这个数字的第几位digitIdx等于(n - 1) % bits
- 计算出digitIdx位的数字
  
  ### 以11为例:
  11 - 9 = 2 
  
  (2 - 1) / 2 = 0
  
  (2 - 1) % 2 = 1
  
  也就是说第11位数字是位数是2的第一个数字的第二位，即是0

## 代码

```go
package leetcode

import "math"

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	bits := 1
	for n > 9*int(math.Pow10(bits-1))*bits {
		n -= 9 * int(math.Pow10(bits-1)) * bits
		bits++
	}
	idx := n - 1
	start := int(math.Pow10(bits - 1))
	num := start + idx/bits
	digitIdx := idx % bits
	return num / int(math.Pow10(bits-digitIdx-1)) % 10
}
```