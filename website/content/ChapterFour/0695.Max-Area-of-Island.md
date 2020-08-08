# [695. Max Area of Island](https://leetcode.com/problems/max-area-of-island/)



## 题目

Given a non-empty 2D array `grid` of 0's and 1's, an **island** is a group of `1`'s (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

**Example 1**:

```
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
```

Given the above grid, return`6`. Note the answer is not 11, because the island must be connected 4-directionally.

**Example 2**:

```
[[0,0,0,0,0,0,0,0]]
```

Given the above grid, return`0`.

**Note**: The length of each dimension in the given `grid` does not exceed 50.

## 题目大意

给定一个包含了一些 0 和 1 的非空二维数组 grid 。一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

## 解题思路

- 给出一个地图，要求计算上面岛屿的面积。注意岛屿的定义是四周都是海(为 0 的点)，如果土地(为 1 的点)靠在地图边缘，不能算是岛屿。
- 这一题和第 200 题，第 1254 题解题思路是一致的。DPS 深搜。这不过这一题需要多处理 2 件事情，一个是注意靠边缘的岛屿不能计算在内，二是动态维护岛屿的最大面积。

## 代码

```go
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				continue
			}
			area := areaOfIsland(grid, i, j)
			if area > res {
				res = area
			}
		}
	}
	return res
}

func areaOfIsland(grid [][]int, x, y int) int {
	if !isInGrid(grid, x, y) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	total := 1
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		total += areaOfIsland(grid, nx, ny)
	}
	return total
}
```