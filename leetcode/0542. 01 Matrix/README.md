# [542. 01 Matrix](https://leetcode.com/problems/01-matrix/)


## 题目

Given a matrix consists of 0 and 1, find the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.

**Example 1:**

    Input:
    [[0,0,0],
     [0,1,0],
     [0,0,0]]
    
    Output:
    [[0,0,0],
     [0,1,0],
     [0,0,0]]

**Example 2:**

    Input:
    [[0,0,0],
     [0,1,0],
     [1,1,1]]
    
    Output:
    [[0,0,0],
     [0,1,0],
     [1,2,1]]

**Note:**

1. The number of elements of the given matrix will not exceed 10,000.
2. There are at least one 0 in the given matrix.
3. The cells are adjacent in only four directions: up, down, left and right.


## 题目大意

给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。两个相邻元素间的距离为 1 。


## 解题思路


- 给出一个二维数组，数组里面只有 0 和 1 。要求计算每个 1 距离最近的 0 的距离。
- 这一题有 3 种解法，第一种解法最容易想到，BFS。先预处理一下棋盘，将每个 0 都处理为 -1 。将 1 都处理为 0 。将每个 -1 (即原棋盘的 0)都入队，每次出队都将四周的 4 个位置都入队。这就想一颗石头扔进了湖里，一圈一圈的波纹荡开，每一圈都是一层。由于棋盘被我们初始化了，所有为 -1 的都是原来为 0 的，所以波纹扫过来不需要处理这些 -1 的点。棋盘上为  0 的点都是原来为 1 的点，这些点在波纹扫过来的时候就需要赋值更新 level。当下次波纹再次扫到原来为 1 的点的时候，由于它已经被第一次到的波纹更新了值，所以这次不用再更新了。(第一次波纹到的时候一定是最短的)
- 第二种解法是 DFS。先预处理，把周围没有 0 的 1 都重置为最大值。当周围有 0 的 1，距离 0 的位置都是 1，这些点是不需要动的，需要更新的点恰恰应该是那些周围没有 0 的点。当递归的步数 val 比点的值小(这也就是为什么会先把 1 更新成最大值的原因)的时候，不断更新它。
- 第三种解法是 DP。由于有 4 个方向，每次处理 2 个方向，可以降低时间复杂度。第一次循环从上到下，从左到右遍历，先处理上边和左边，第二次循环从下到上，从右到左遍历，再处理右边和下边。
