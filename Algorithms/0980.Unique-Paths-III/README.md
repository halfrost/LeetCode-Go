# [980. Unique Paths III](https://leetcode.com/problems/unique-paths-iii/)


## 题目

On a 2-dimensional `grid`, there are 4 types of squares:

- `1` represents the starting square. There is exactly one starting square.
- `2` represents the ending square. There is exactly one ending square.
- `0` represents empty squares we can walk over.
- `-1` represents obstacles that we cannot walk over.

Return the number of 4-directional walks from the starting square to the ending square, that **walk over every non-obstacle square exactly once**.

**Example 1:**

    Input: [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
    Output: 2
    Explanation: We have the following two paths: 
    1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
    2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)

**Example 2:**

    Input: [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
    Output: 4
    Explanation: We have the following four paths: 
    1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
    2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
    3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
    4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)

**Example 3:**

    Input: [[0,1],[2,0]]
    Output: 0
    Explanation: 
    There is no path that walks over every empty square exactly once.
    Note that the starting and ending square can be anywhere in the grid.

**Note:**

1. `1 <= grid.length * grid[0].length <= 20`


## 题目大意

在二维网格 grid 上，有 4 种类型的方格：

- 1 表示起始方格。且只有一个起始方格。
- 2 表示结束方格，且只有一个结束方格。
- 0 表示我们可以走过的空方格。
- -1 表示我们无法跨越的障碍。

返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目，**每一个无障碍方格都要通过一次**。



## 解题思路


- 这一题也可以按照第 79 题的思路来做。题目要求输出地图中从起点到终点的路径条数。注意路径要求必须走满所有空白的格子。
- 唯一需要注意的一点是，空白的格子并不是最后走的总步数，`总步数 = 空白格子数 + 1`，因为要走到终点，走到终点也算一步。
