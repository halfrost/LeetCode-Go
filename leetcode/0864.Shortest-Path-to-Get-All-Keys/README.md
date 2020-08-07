# [864. Shortest Path to Get All Keys](https://leetcode.com/problems/shortest-path-to-get-all-keys/)


## 题目:

We are given a 2-dimensional `grid`. `"."` is an empty cell, `"#"` is a wall, `"@"` is the starting point, (`"a"`, `"b"`, ...) are keys, and (`"A"`, `"B"`, ...) are locks.

We start at the starting point, and one move consists of walking one space in one of the 4 cardinal directions. We cannot walk outside the grid, or walk into a wall. If we walk over a key, we pick it up. We can't walk over a lock unless we have the corresponding key.

For some 1 <= K <= 6, there is exactly one lowercase and one uppercase letter of the first `K` letters of the English alphabet in the grid. This means that there is exactly one key for each lock, and one lock for each key; and also that the letters used to represent the keys and locks were chosen in the same order as the English alphabet.

Return the lowest number of moves to acquire all keys. If it's impossible, return `-1`.

**Example 1:**

    Input: ["@.a.#","###.#","b.A.B"]
    Output: 8

**Example 2:**

    Input: ["@..aA","..B#.","....b"]
    Output: 6

**Note:**

1. `1 <= grid.length <= 30`
2. `1 <= grid[0].length <= 30`
3. `grid[i][j]` contains only `'.'`, `'#'`, `'@'`, `'a'-'f'` and `'A'-'F'`
4. The number of keys is in `[1, 6]`. Each key has a different letter and opens exactly one lock.


## 题目大意

给定一个二维网格 grid。 "." 代表一个空房间， "#" 代表一堵墙， "@" 是起点，（"a", "b", ...）代表钥匙，（"A", "B", ...）代表锁。

我们从起点开始出发，一次移动是指向四个基本方向之一行走一个单位空间。我们不能在网格外面行走，也无法穿过一堵墙。如果途经一个钥匙，我们就把它捡起来。除非我们手里有对应的钥匙，否则无法通过锁。

假设 K 为钥匙/锁的个数，且满足 1 <= K <= 6，字母表中的前 K 个字母在网格中都有自己对应的一个小写和一个大写字母。换言之，每个锁有唯一对应的钥匙，每个钥匙也有唯一对应的锁。另外，代表钥匙和锁的字母互为大小写并按字母顺序排列。

返回获取所有钥匙所需要的移动的最少次数。如果无法获取所有钥匙，返回 -1 。

提示：

1. 1 <= grid.length <= 30
2. 1 <= grid[0].length <= 30
3. grid[i][j] 只含有 '.', '#', '@', 'a'-'f' 以及 'A'-'F'
4. 钥匙的数目范围是 [1, 6]，每个钥匙都对应一个不同的字母，正好打开一个对应的锁。


## 解题思路


- 给出一个地图，在图中有钥匙和锁，当锁在没有钥匙的时候不能通行，问从起点 @ 开始，到最终获得所有钥匙，最短需要走多少步。
- 这一题可以用 BFS 来解答。由于钥匙的种类比较多，所以 visited 数组需要 3 个维度，一个是 x 坐标，一个是 y 坐标，最后一个是当前获取钥匙的状态。每把钥匙都有获取了和没有获取两种状态，题目中说最多有 6 把钥匙，那么排列组合最多是 2^6 = 64 种状态。用一个十进制数的二进制位来压缩这些状态，二进制位分别来表示这些钥匙是否已经获取了。既然钥匙的状态可以压缩，其实 x 和 y 的坐标也可以一并压缩到这个数中。BFS 中存的数字是坐标 + 钥匙状态的状态。在 BFS 遍历的过程中，用 visited 数组来过滤遍历过的情况，来保证走的路是最短的。其他的情况无非是判断锁的状态，是否能通过，判断钥匙获取状态。
- 这一题不知道是否能用 DFS 来解答。我实现了一版，但是在 18 / 35 这组 case 上超时了，具体 case 见测试文件第一个 case。
