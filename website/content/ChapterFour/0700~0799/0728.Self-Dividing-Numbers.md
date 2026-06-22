# [728. Self Dividing Numbers](https://leetcode.com/problems/self-dividing-numbers/)

## 题目

A self-dividing number is a number that is divisible by every digit it contains.

- For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.

A self-dividing number is not allowed to contain the digit zero.

Given two integers left and right, return a list of all the self-dividing numbers in the range [left, right].

**Example 1:**

    Input: left = 1, right = 22
    Output: [1,2,3,4,5,6,7,8,9,11,12,15,22]

**Example 2:**

    Input: left = 47, right = 85
    Output: [48,55,66,77]

**Constraints:**

- 1 <= left <= right <= 10000

## 题目大意

自除数是指可以被它包含的每一位数整除的数。

- 例如，128 是一个 自除数 ，因为 128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。

自除数 不允许包含 0 。

给定两个整数 left 和 right ，返回一个列表，列表的元素是范围 [left, right] 内所有的 自除数 。

## 解题思路

- 模拟计算

# 代码

```go
package leetcode

func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for num := left; num <= right; num++ {
		if selfDividingNum(num) {
			ans = append(ans, num)
		}
	}
	return ans
}

func selfDividingNum(num int) bool {
	for d := num; d > 0; d = d / 10 {
		reminder := d % 10
		if reminder == 0 {
			return false
		}
		if num%reminder != 0 {
			return false
		}
	}
	return true
}
```
