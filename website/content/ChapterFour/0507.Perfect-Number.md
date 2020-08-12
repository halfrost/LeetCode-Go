# [507. Perfect Number](https://leetcode.com/problems/perfect-number/)



## 题目

We define the Perfect Number is a **positive** integer that is equal to the sum of all its **positive** divisors except itself.

Now, given an

**integer**

n, write a function that returns true when it is a perfect number and false when it is not.

**Example**:

```
Input: 28
Output: True
Explanation: 28 = 1 + 2 + 4 + 7 + 14
```

**Note**: The input number **n** will not exceed 100,000,000. (1e8)

## 题目大意

对于一个 正整数，如果它和除了它自身以外的所有正因子之和相等，我们称它为“完美数”。给定一个 整数 n， 如果他是完美数，返回 True，否则返回 False

## 解题思路

- 给定一个整数，要求判断这个数是不是完美数。整数的取值范围小于 1e8 。
- 简单题。按照题意描述，先获取这个整数的所有正因子，如果正因子的和等于原来这个数，那么它就是完美数。
- 这一题也可以打表，1e8 以下的完美数其实并不多，就 5 个。

## 代码

```go

package leetcode

import "math"

// 方法一
func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum, bound := 1, int(math.Sqrt(float64(num)))+1
	for i := 2; i < bound; i++ {
		if num%i != 0 {
			continue
		}
		corrDiv := num / i
		sum += corrDiv + i
	}
	return sum == num
}

// 方法二 打表
func checkPerfectNumber_(num int) bool {
	return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
}

```