# [1654. Minimum Jumps to Reach Home](https://leetcode.com/problems/minimum-jumps-to-reach-home/)


## 题目

A certain bug's home is on the x-axis at position `x`. Help them get there from position `0`.

The bug jumps according to the following rules:

- It can jump exactly `a` positions **forward** (to the right).
- It can jump exactly `b` positions **backward** (to the left).
- It cannot jump backward twice in a row.
- It cannot jump to any `forbidden` positions.

The bug may jump forward **beyond** its home, but it **cannot jump** to positions numbered with **negative** integers.

Given an array of integers `forbidden`, where `forbidden[i]` means that the bug cannot jump to the position `forbidden[i]`, and integers `a`, `b`, and `x`, return *the minimum number of jumps needed for the bug to reach its home*. If there is no possible sequence of jumps that lands the bug on position `x`, return `1.`

**Example 1:**

```
Input: forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
Output: 3
Explanation: 3 jumps forward (0 -> 3 -> 6 -> 9) will get the bug home.
```

**Example 2:**

```
Input: forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
Output: -1
```

**Example 3:**

```
Input: forbidden = [1,6,2,14,5,17,4], a = 16, b = 9, x = 7
Output: 2
Explanation: One jump forward (0 -> 16) then one jump backward (16 -> 7) will get the bug home.

```

**Constraints:**

- `1 <= forbidden.length <= 1000`
- `1 <= a, b, forbidden[i] <= 2000`
- `0 <= x <= 2000`
- All the elements in `forbidden` are distinct.
- Position `x` is not forbidden.

## 题目大意

有一只跳蚤的家在数轴上的位置 x 处。请你帮助它从位置 0 出发，到达它的家。

跳蚤跳跃的规则如下：

- 它可以 往前 跳恰好 a 个位置（即往右跳）。
- 它可以 往后 跳恰好 b 个位置（即往左跳）。
- 它不能 连续 往后跳 2 次。
- 它不能跳到任何 forbidden 数组中的位置。

跳蚤可以往前跳 超过 它的家的位置，但是它 不能跳到负整数 的位置。给你一个整数数组 forbidden ，其中 forbidden[i] 是跳蚤不能跳到的位置，同时给你整数 a， b 和 x ，请你返回跳蚤到家的最少跳跃次数。如果没有恰好到达 x 的可行方案，请你返回 -1 。

## 解题思路

- 给出坐标 x ，可以往前跳的步长 a，往后跳的步长 b。要求输出能跳回家的最少跳跃次数。
- 求最少跳跃次数，思路用 BFS 求解，最先到达坐标 x 的方案即是最少跳跃次数。对`forbidden` 的处理是把记忆化数组里面把他们标记为 true。禁止连续往后跳 2 次的限制，要求我们在 BFS 入队的时候再记录一下跳跃方向，每次往后跳的时候判断前一跳是否是往后跳，如果是往后跳，此次就不能往后跳了。

## 代码

```go
package leetcode

func minimumJumps(forbidden []int, a int, b int, x int) int {
	visited := make([]bool, 6000)
	for i := range forbidden {
		visited[forbidden[i]] = true
	}
	queue, res := [][2]int{{0, 0}}, -1
	for len(queue) > 0 {
		length := len(queue)
		res++
		for i := 0; i < length; i++ {
			cur, isBack := queue[i][0], queue[i][1]
			if cur == x {
				return res
			}
			if isBack == 0 && cur-b > 0 && !visited[cur-b] {
				visited[cur-b] = true
				queue = append(queue, [2]int{cur - b, 1})
			}
			if cur+a < len(visited) && !visited[cur+a] {
				visited[cur+a] = true
				queue = append(queue, [2]int{cur + a, 0})
			}
		}
		queue = queue[length:]
	}
	return -1
}
```