# [1486. XOR Operation in an Array](https://leetcode.com/problems/xor-operation-in-an-array/)


## 题目

Given an integer `n` and an integer `start`.

Define an array `nums` where `nums[i] = start + 2*i` (0-indexed) and `n == nums.length`.

Return the bitwise XOR of all elements of `nums`.

**Example 1:**

```
Input: n = 5, start = 0
Output: 8
Explanation:Array nums is equal to [0, 2, 4, 6, 8] where (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8.
Where "^" corresponds to bitwise XOR operator.
```

**Example 2:**

```
Input: n = 4, start = 3
Output: 8
Explanation:Array nums is equal to [3, 5, 7, 9] where (3 ^ 5 ^ 7 ^ 9) = 8.
```

**Example 3:**

```
Input: n = 1, start = 7
Output: 7
```

**Example 4:**

```
Input: n = 10, start = 5
Output: 2
```

**Constraints:**

- `1 <= n <= 1000`
- `0 <= start <= 1000`
- `n == nums.length`

## 题目大意

给你两个整数，n 和 start 。数组 nums 定义为：nums[i] = start + 2*i（下标从 0 开始）且 n == nums.length 。请返回 nums 中所有元素按位异或（XOR）后得到的结果。

## 解题思路

- 简单题。按照题意，一层循环依次累积异或数组中每个元素。

## 代码

```go
package leetcode

func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + 2*i
	}
	return res
}
```