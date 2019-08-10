# [778. Swim in Rising Water](https://leetcode.com/problems/swim-in-rising-water/)


## 题目:

On an N x N `grid`, each square `grid[i][j]` represents the elevation at that point `(i,j)`.

Now rain starts to fall. At time `t`, the depth of the water everywhere is `t`. You can swim from a square to another 4-directionally adjacent square if and only if the elevation of both squares individually are at most `t`. You can swim infinite distance in zero time. Of course, you must stay within the boundaries of the grid during your swim.

You start at the top left square `(0, 0)`. What is the least time until you can reach the bottom right square `(N-1, N-1)`?

**Example 1:**

    Input: [[0,2],[1,3]]
    Output: 3
    Explanation:
    At time 0, you are in grid location (0, 0).
    You cannot go anywhere else because 4-directionally adjacent neighbors have a higher elevation than t = 0.
    
    You cannot reach point (1, 1) until time 3.
    When the depth of water is 3, we can swim anywhere inside the grid.

**Example 2:**

    Input: [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
    Output: 16
    Explanation:
     0  1  2  3  4
    24 23 22 21  5
    12 13 14 15 16
    11 17 18 19 20
    10  9  8  7  6
    
    The final route is marked in bold.
    We need to wait until time 16 so that (0, 0) and (4, 4) are connected.

**Note:**

1. `2 <= N <= 50`.
2. grid[i][j] is a permutation of [0, ..., N*N - 1].

## 题目大意


在一个 N x N 的坐标方格 grid 中，每一个方格的值 grid[i][j] 表示在位置 (i,j) 的平台高度。现在开始下雨了。当时间为 t 时，此时雨水导致水池中任意位置的水位为 t 。你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两个平台。假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。

你从坐标方格的左上平台 (0，0) 出发。最少耗时多久你才能到达坐标方格的右下平台 (N-1, N-1)？

提示:

- 2 <= N <= 50.
- grid[i][j] 位于区间 [0, ..., N*N - 1] 内。


## 解题思路

- 给出一个 grid[i][j] 方格，每个格子里面表示游泳池里面平台的高度。t 时刻，游泳池中的水的高度是 t。只有水的高度到达了平台的高度以后才能游过去。问从 (0,0) 开始，最短多长时间能到达 (N-1, N-1) 。
- 这一题有多种解法。第一种解题思路是利用 DFS + 二分。DFS 是用来遍历是否可达。利用时间(即当前水淹过的高度)来判断是否能到达终点 (N-1, N-1) 点。二分用来搜索最终结果的时间。为什么会考虑用二分加速呢？原因是：时间从 0 - max 依次递增。max 是游泳池最高的平台高度。当时间从 0 增加到 max 以后，肯定能到达终点 (N-1, N-1) 点，因为水比所有平台都要高了。想快速找到一个时间 t 能使得 (0,0) 点和 (N-1, N-1) 点之间连通，那么就想到用二分加速了。判断是否取中值的条件是 (0,0) 点和 (N-1, N-1) 点之间是否连通。
- 第二种解题思路是并查集。只要是 (0,0) 点和 (N-1, N-1) 点没有连通，即不能游到终点，那么就开始 `union()` 操作，由于起点是 (0,0)，所以向右边 `i + 1` 和向下边 `j + 1` 开始尝试。每尝试完一轮，时间会加 1 秒，即高度会加一。直到 (0,0) 点和 (N-1, N-1) 点刚好连通，那么这个时间点就是最终要求的。