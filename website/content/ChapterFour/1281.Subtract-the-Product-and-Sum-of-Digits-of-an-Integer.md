# [1281. Subtract the Product and Sum of Digits of an Integer](https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/)



## 题目

Given an integer number `n`, return the difference between the product of its digits and the sum of its digits.

**Example 1**:

```
Input: n = 234
Output: 15 
Explanation: 
Product of digits = 2 * 3 * 4 = 24 
Sum of digits = 2 + 3 + 4 = 9 
Result = 24 - 9 = 15
```

**Example 2**:

```
Input: n = 4421
Output: 21
Explanation: 
Product of digits = 4 * 4 * 2 * 1 = 32 
Sum of digits = 4 + 4 + 2 + 1 = 11 
Result = 32 - 11 = 21
```

**Constraints**:

- `1 <= n <= 10^5`

## 题目大意

给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。

提示：

- 1 <= n <= 10^5

## 解题思路

- 给出一个数，计算这个数每位数字乘积减去每位数字累加的差值。
- 简单题，按照题意输入输出即可。

## 代码

```go
func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for ; n > 0; n /= 10 {
		sum += n % 10
		product *= n % 10
	}
	return product - sum
}
```