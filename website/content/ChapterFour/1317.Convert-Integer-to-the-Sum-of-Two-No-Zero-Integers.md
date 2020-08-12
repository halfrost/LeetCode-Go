# [1317. Convert Integer to the Sum of Two No-Zero Integers](https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/)


## 题目

Given an integer `n`. No-Zero integer is a positive integer which **doesn't contain any 0** in its decimal representation.

Return *a list of two integers* `[A, B]` where:

- `A` and `B` are No-Zero integers.
- `A + B = n`

It's guarateed that there is at least one valid solution. If there are many valid solutions you can return any of them.

**Example 1**:

```
Input: n = 2
Output: [1,1]
Explanation: A = 1, B = 1. A + B = n and both A and B don't contain any 0 in their decimal representation.
```

**Example 2**:

```
Input: n = 11
Output: [2,9]
```

**Example 3**:

```
Input: n = 10000
Output: [1,9999]
```

**Example 4**:

```
Input: n = 69
Output: [1,68]
```

**Example 5**:

```
Input: n = 1010
Output: [11,999]
```

**Constraints**:

- `2 <= n <= 10^4`

## 题目大意

「无零整数」是十进制表示中 不含任何 0 的正整数。给你一个整数 n，请你返回一个 由两个整数组成的列表 [A, B]，满足：

- A 和 B 都是无零整数
- A + B = n

题目数据保证至少有一个有效的解决方案。如果存在多个有效解决方案，你可以返回其中任意一个。

## 解题思路

- 给定一个整数 n，要求把它分解为 2 个十进制位中不含 0 的正整数且这两个正整数之和为 n。
- 简单题。在 [1, n/2] 区间内搜索，只要有一组满足条件的解就 break。题目保证了至少有一组解，并且多组解返回任意一组即可。

## 代码

```go

package leetcode

func getNoZeroIntegers(n int) []int {
	noZeroPair := []int{}
	for i := 1; i <= n/2; i++ {
		if isNoZero(i) && isNoZero(n-i) {
			noZeroPair = append(noZeroPair, []int{i, n - i}...)
			break
		}
	}
	return noZeroPair
}

func isNoZero(n int) bool {
	for n != 0 {
		if n%10 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

```