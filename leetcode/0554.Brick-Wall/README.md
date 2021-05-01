# [554. Brick Wall](https://leetcode.com/problems/brick-wall/)

## 题目

There is a rectangular brick wall in front of you with `n` rows of bricks. The `ith` row has some number of bricks each of the same height (i.e., one unit) but they can be of different widths. The total width of each row is the same.

Draw a vertical line from the top to the bottom and cross the least bricks. If your line goes through the edge of a brick, then the brick is not considered as crossed. You cannot draw a line just along one of the two vertical edges of the wall, in which case the line will obviously cross no bricks.

Given the 2D array `wall` that contains the information about the wall, return *the minimum number of crossed bricks after drawing such a vertical line*.

**Example 1:**

![https://assets.leetcode.com/uploads/2021/04/24/cutwall-grid.jpg](https://assets.leetcode.com/uploads/2021/04/24/cutwall-grid.jpg)

```
Input: wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
Output: 2

```

**Example 2:**

```
Input: wall = [[1],[1],[1]]
Output: 3

```

**Constraints:**

- `n == wall.length`
- `1 <= n <= 10^4`
- `1 <= wall[i].length <= 10^4`
- `1 <= sum(wall[i].length) <= 2 * 10^4`
- `sum(wall[i])` is the same for each row `i`.
- `1 <= wall[i][j] <= 2^31 - 1`

## 题目大意

你的面前有一堵矩形的、由 n 行砖块组成的砖墙。这些砖块高度相同（也就是一个单位高）但是宽度不同。每一行砖块的宽度之和应该相等。你现在要画一条 自顶向下 的、穿过 最少 砖块的垂线。如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。给你一个二维数组 wall ，该数组包含这堵墙的相关信息。其中，wall[i] 是一个代表从左至右每块砖的宽度的数组。你需要找出怎样画才能使这条线 穿过的砖块数量最少 ，并且返回 穿过的砖块数量 。

## 解题思路

- 既然穿过砖块中缝不算穿过砖块，那么穿过最少砖块数量一定是穿过很多中缝。按行遍历每一行的砖块，累加每行砖块宽度，将每行砖块“缝”的坐标存在 map 中。最后取出 map 中出现频次最高的缝，即为铅垂线要穿过的地方。墙高减去缝出现的频次，剩下的即为穿过砖块的数量。

## 代码

```go
package leetcode

func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	for _, row := range wall {
		sum := 0
		for i := 0; i < len(row)-1; i++ {
			sum += row[i]
			m[sum]++
		}
	}
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return len(wall) - max
}
```