# [55. Jump Game](https://leetcode.com/problems/jump-game/)


## 题目

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

**Example 1**:

```
Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

**Example 2**:

```
Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
```

## 题目大意

给定一个非负整数数组，最初位于数组的第一个位置。数组中的每个元素代表在该位置可以跳跃的最大长度。判断是否能够到达最后一个位置。

## 解题思路

- 给出一个非负数组，要求判断从数组 0 下标开始，能否到达数组最后一个位置。
- 这一题比较简单。如果某一个作为 `起跳点` 的格子可以跳跃的距离是 `n`，那么表示后面 `n` 个格子都可以作为 `起跳点`。可以对每一个能作为 `起跳点` 的格子都尝试跳一次，把 `能跳到最远的距离maxJump` 不断更新。如果可以一直跳到最后，就成功了。如果中间有一个点比 `maxJump` 还要大，说明在这个点和 maxJump 中间连不上了，有些点不能到达最后一个位置。

## 代码

```go
func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+v)
	}
	return true
}
```