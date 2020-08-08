# [1304. Find N Unique Integers Sum up to Zero](https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/)



## 题目

Given an integer `n`, return **any** array containing `n` **unique** integers such that they add up to 0.

**Example 1**:

```
Input: n = 5
Output: [-7,-1,1,3,4]
Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].
```

**Example 2**:

```
Input: n = 3
Output: [-1,0,1]
```

**Example 3**:

```
Input: n = 1
Output: [0]
```

**Constraints**:

- `1 <= n <= 1000`

## 题目大意

给你一个整数 n，请你返回 任意 一个由 n 个 各不相同 的整数组成的数组，并且这 n 个数相加和为 0 。

提示：

- 1 <= n <= 1000

## 解题思路

- 给出一个数 n，输出一个有 n 个数的数组，里面元素之和为 0 。
- 简单题，简单循环即可。

## 代码

```go
func sumZero(n int) []int {
	res, left, right, start := make([]int, n), 0, n-1, 1
	for left < right {
		res[left] = start
		res[right] = -start
		start++
		left = left + 1
		right = right - 1
	}
	return res
}
```