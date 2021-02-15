# [1732. Find the Highest Altitude](https://leetcode.com/problems/find-the-highest-altitude/)


## 题目

There is a biker going on a road trip. The road trip consists of `n + 1` points at different altitudes. The biker starts his trip on point `0` with altitude equal `0`.

You are given an integer array `gain` of length `n` where `gain[i]` is the **net gain in altitude** between points `i` and `i + 1` for all (`0 <= i < n)`. Return *the **highest altitude** of a point.*

**Example 1:**

```
Input: gain = [-5,1,5,0,-7]
Output: 1
Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
```

**Example 2:**

```
Input: gain = [-4,-3,-2,-1,4,3,2]
Output: 0
Explanation: The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.
```

**Constraints:**

- `n == gain.length`
- `1 <= n <= 100`
- `100 <= gain[i] <= 100`

## 题目大意

有一个自行车手打算进行一场公路骑行，这条路线总共由 n + 1 个不同海拔的点组成。自行车手从海拔为 0 的点 0 开始骑行。给你一个长度为 n 的整数数组 gain ，其中 gain[i] 是点 i 和点 i + 1 的 净海拔高度差（0 <= i < n）。请你返回 最高点的海拔 。

## 解题思路

- 简单题。循环数组依次从第一个海拔点开始还原每个海拔点，动态记录最大高度。循环结束输出最大高度即可。

## 代码

```go
package leetcode

func largestAltitude(gain []int) int {
	max, height := 0, 0
	for _, g := range gain {
		height += g
		if height > max {
			max = height
		}
	}
	return max
}
```