# [1293. Shortest Path in a Grid with Obstacles Elimination](https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/)



## 题目

You are given an m x n integer matrix grid where each cell is either 0 (empty) or 1 (obstacle). You can move up, down, left, or right from and to an empty cell in one step.

Return the minimum number of steps to walk from the upper left corner (0, 0) to the lower right corner (m - 1, n - 1) given that you can eliminate at most k obstacles. If it is not possible to find such walk return -1.



Example 1:


![](https://assets.leetcode.com/uploads/2021/09/30/short1-grid.jpg)
 

```
Input: grid = [[0,0,0],[1,1,0],[0,0,0],[0,1,1],[0,0,0]], k = 1
Output: 6
Explanation: 
The shortest path without eliminating any obstacle is 10.
The shortest path with one obstacle elimination at position (3,2) is 6. Such path is (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3,2) -> (4,2).
```

Example 2:

![](https://assets.leetcode.com/uploads/2021/09/30/short2-grid.jpg)

```
Input: grid = [[0,1,1],[1,1,1],[1,0,0]], k = 1
Output: -1
Explanation: We need to eliminate at least two obstacles to find such a walk.
```

Constraints:

- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 40
- 1 <= k <= m * n
- grid[i][j] is either 0 or 1.
- grid[0][0] == grid[m - 1][n - 1] == 0



## 题目大意

给你一个 m * n 的网格，其中每个单元格不是 0（空）就是 1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。

如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。如果找不到这样的路径，则返回 -1 。


## 解题思路

使用 BFS 遍历棋盘。这题比普通可达性问题多了一个障碍物的限制。这个也不难。每个点往周边四个方向扩展的时候，如果遇到障碍物，先算上这个障碍物，障碍物累积总个数小于 K 的时候，从障碍物的这个格子继续开始遍历。如果没有遇到障碍物，判断当前累积障碍物个数是否已经小于 K 个，如果小于 K 便继续遍历。如果大于 K，便终止此轮遍历。

## 代码

```go
var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type pos struct {
	x, y     int
	obstacle int
	step     int
}

func shortestPath(grid [][]int, k int) int {
	queue, m, n := []pos{}, len(grid), len(grid[0])
	visitor := make([][][]int, m)
	if len(grid) == 1 && len(grid[0]) == 1 {
		return 0
	}
	for i := 0; i < m; i++ {
		visitor[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			visitor[i][j] = make([]int, k+1)
		}
	}
	visitor[0][0][0] = 1
	queue = append(queue, pos{x: 0, y: 0, obstacle: 0, step: 0})
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			for i := 0; i < len(dir); i++ {
				newX := node.x + dir[i][0]
				newY := node.y + dir[i][1]
				if newX == m-1 && newY == n-1 {
					if node.obstacle != 0 {
						if node.obstacle <= k {
							return node.step + 1
						} else {
							continue
						}
					}
					return node.step + 1
				}
				if isInBoard(grid, newX, newY) {
					if grid[newX][newY] == 1 {
						if node.obstacle+1 <= k && visitor[newX][newY][node.obstacle+1] != 1 {
							queue = append(queue, pos{x: newX, y: newY, obstacle: node.obstacle + 1, step: node.step + 1})
							visitor[newX][newY][node.obstacle+1] = 1
						}
					} else {
						if node.obstacle <= k && visitor[newX][newY][node.obstacle] != 1 {
							queue = append(queue, pos{x: newX, y: newY, obstacle: node.obstacle, step: node.step + 1})
							visitor[newX][newY][node.obstacle] = 1
						}
					}

				}
			}
		}
	}
	return -1
}

func isInBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
```