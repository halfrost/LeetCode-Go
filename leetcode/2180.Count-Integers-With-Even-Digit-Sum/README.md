# [2180. Count Integers With Even Digit Sum](https://leetcode.com/problems/count-integers-with-even-digit-sum/)


## 题目

Given a positive integer `num`, return *the number of positive integers **less than or equal to*** `num` *whose digit sums are **even***.

The **digit sum** of a positive integer is the sum of all its digits.

**Example 1:**

```
Input: num = 4
Output: 2
Explanation:
The only integers less than or equal to 4 whose digit sums are even are 2 and 4.

```

**Example 2:**

```
Input: num = 30
Output: 14
Explanation:
The 14 integers less than or equal to 30 whose digit sums are even are
2, 4, 6, 8, 11, 13, 15, 17, 19, 20, 22, 24, 26, and 28.

```

**Constraints:**

- `1 <= num <= 1000`

## 题目大意

给你一个正整数 num ，请你统计并返回 小于或等于 num 且各位数字之和为 偶数 的正整数的数目。

正整数的 各位数字之和 是其所有位上的对应数字相加的结果。

## 解题思路

- 简单题。依照题意，计算每个数的各位数字之和，如何和为偶数，则统计结果加一。最后输出统计结果即可。

## 代码

```go
package leetcode

func countEven(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if addSum(i)%2 == 0 {
			count++
		}
	}
	return count
}

func addSum(num int) int {
	sum := 0
	tmp := num
	for tmp != 0 {
		sum += tmp % 10
		tmp = tmp / 10
	}
	return sum
}
```