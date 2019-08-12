# [815. Bus Routes](https://leetcode.com/problems/bus-routes/)


## 题目:

We have a list of bus routes. Each `routes[i]` is a bus route that the i-th bus repeats forever. For example if `routes[0] = [1, 5, 7]`, this means that the first bus (0-th indexed) travels in the sequence 1->5->7->1->5->7->1->... forever.

We start at bus stop `S` (initially not on a bus), and we want to go to bus stop `T`. Travelling by buses only, what is the least number of buses we must take to reach our destination? Return -1 if it is not possible.

    Example:
    Input: 
    routes = [[1, 2, 7], [3, 6, 7]]
    S = 1
    T = 6
    Output: 2
    Explanation: 
    The best strategy is take the first bus to the bus stop 7, then take the second bus to the bus stop 6.

**Note:**

- `1 <= routes.length <= 500`.
- `1 <= routes[i].length <= 500`.
- `0 <= routes[i][j] < 10 ^ 6`.


## 题目大意

我们有一系列公交路线。每一条路线 routes[i] 上都有一辆公交车在上面循环行驶。例如，有一条路线 routes[0] = [1, 5, 7]，表示第一辆 (下标为0) 公交车会一直按照 1->5->7->1->5->7->1->... 的车站路线行驶。假设我们从 S 车站开始（初始时不在公交车上），要去往 T 站。 期间仅可乘坐公交车，求出最少乘坐的公交车数量。返回 -1 表示不可能到达终点车站。


说明:

- 1 <= routes.length <= 500.
- 1 <= routes[i].length <= 500.
- 0 <= routes[i][j] < 10 ^ 6.


## 解题思路

- 给出一些公交路线，公交路径代表经过的哪些站。现在给出起点和终点站，问最少需要换多少辆公交车才能从起点到终点？
- 这一题可以转换成图论的问题，将每个站台看成顶点，公交路径看成每个顶点的边。同一个公交的边染色相同。题目即可转化为从顶点 S 到顶点 T 需要经过最少多少条不同的染色边。用 BFS 即可轻松解决。从起点 S 开始，不断的扩展它能到达的站点。用 visited 数组防止放入已经可达的站点引起的环。用 map 存储站点和公交车的映射关系(即某个站点可以由哪些公交车到达)，BFS 的过程中可以用这个映射关系，拿到公交车的其他站点信息，从而扩张队列里面的可达站点。一旦扩展出现了终点 T，就可以返回结果了。
