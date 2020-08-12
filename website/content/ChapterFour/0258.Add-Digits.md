# [258. Add Digits](https://leetcode.com/problems/add-digits/)


## 题目

Given a non-negative integer `num`, repeatedly add all its digits until the result has only one digit.

**Example**:

```
Input: 38
Output: 2 
Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2. 
             Since 2 has only one digit, return it.
```

**Follow up**: Could you do it without any loop/recursion in O(1) runtime?

## 题目大意

给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。


## 解题思路

- 给定一个非负整数，反复加各个位上的数，直到结果为一位数为止，最后输出这一位数。
- 简单题。按照题意循环累加即可。

## 代码

```go

package leetcode

func addDigits(num int) int {
	for num > 9 {
		cur := 0
		for num != 0 {
			cur += num % 10
			num /= 10
		}
		num = cur
	}
	return num
}

```