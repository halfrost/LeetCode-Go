# [1020. Number of Enclaves](https://leetcode.com/problems/number-of-enclaves/)



## 题目

Given a 2D array `A`, each cell is 0 (representing sea) or 1 (representing land)

A move consists of walking from one land square 4-directionally to another land square, or off the boundary of the grid.

Return the number of land squares in the grid for which we **cannot** walk off the boundary of the grid in any number of moves.

**Example 1**:

```
Input: [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
Output: 3
Explanation: 
There are three 1s that are enclosed by 0s, and one 1 that isn't enclosed because its on the boundary.
```

**Example 2**:

```
Input: [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
Output: 0
Explanation: 
All 1s are either on the boundary or can reach the boundary.
```

**Note**:

1. `1 <= A.length <= 500`
2. `1 <= A[i].length <= 500`
3. `0 <= A[i][j] <= 1`
4. All rows have the same size.

## 题目大意

给出一个二维数组 A，每个单元格为 0（代表海）或 1（代表陆地）。移动是指在陆地上从一个地方走到另一个地方（朝四个方向之一）或离开网格的边界。返回网格中无法在任意次数的移动中离开网格边界的陆地单元格的数量。

提示：

- 1 <= A.length <= 500
- 1 <= A[i].length <= 500
- 0 <= A[i][j] <= 1
- 所有行的大小都相同


## 解题思路

- 给出一个地图，要求输出不和边界连通的 1 的个数。
- 这一题可以用 DFS 也可以用并查集解答。DFS 的思路是深搜的过程中把和边界连通的点都覆盖成 0，最后遍历一遍地图，输出 1 的个数即可。并查集的思路就比较直接了，把能和边界连通的放在一个集合中，剩下的就是不能和边界连通的都在另外一个集合中，输出这个集合里面元素的个数即可。
- 这一题和第 200 题，第 1254 题，第 695 题类似。可以放在一起练习。

## 代码

```go
func numEnclaves(A [][]int) int {
	m, n := len(A), len(A[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				if A[i][j] == 1 {
					dfsNumEnclaves(A, i, j)
				}
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func dfsNumEnclaves(A [][]int, x, y int) {
	if !isInGrid(A, x, y) || A[x][y] == 0 {
		return
	}
	A[x][y] = 0
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		dfsNumEnclaves(A, nx, ny)
	}
}

```