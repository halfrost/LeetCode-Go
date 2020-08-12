# [892. Surface Area of 3D Shapes](https://leetcode.com/problems/surface-area-of-3d-shapes/)


## 题目

On a `N * N` grid, we place some `1 * 1 * 1` cubes.

Each value `v = grid[i][j]` represents a tower of `v` cubes placed on top of grid cell `(i, j)`.

Return the total surface area of the resulting shapes.

**Example 1**:

```
Input: [[2]]
Output: 10
```

**Example 2**:

```
Input: [[1,2],[3,4]]
Output: 34
```

**Example 3**:

```
Input: [[1,0],[0,2]]
Output: 16
```

**Example 4**:

```
Input: [[1,1,1],[1,0,1],[1,1,1]]
Output: 32
```

**Example 5**:

```
Input: [[2,2,2],[2,1,2],[2,2,2]]
Output: 46
```

**Note**:

- `1 <= N <= 50`
- `0 <= grid[i][j] <= 50`

## 题目大意

在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。请你返回最终形体的表面积。


## 解题思路

- 给定一个网格数组，数组里面装的是立方体叠放在所在的单元格，求最终这些叠放的立方体的表面积。
- 简单题。按照题目意思，找到叠放时，重叠的面，然后用总表面积减去这些重叠的面积即为最终答案。

## 代码

```go

package leetcode

func surfaceArea(grid [][]int) int {
	area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			area += grid[i][j]*4 + 2
			// up
			if i > 0 {
				m := min(grid[i][j], grid[i-1][j])
				area -= m
			}
			// down
			if i < len(grid)-1 {
				m := min(grid[i][j], grid[i+1][j])
				area -= m
			}
			// left
			if j > 0 {
				m := min(grid[i][j], grid[i][j-1])
				area -= m
			}
			// right
			if j < len(grid[i])-1 {
				m := min(grid[i][j], grid[i][j+1])
				area -= m
			}
		}
	}
	return area
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

```