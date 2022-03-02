# [1791.Find Center of Star Graph](https://leetcode.com/problems/find-center-of-star-graph/)

## 题目

There is an undirected star graph consisting of n nodes labeled from 1 to n. A star graph is a graph where there is one center node and exactly n - 1 edges that connect the center node with every other node.

You are given a 2D integer array edges where each edges[i] = [ui, vi] indicates that there is an edge between the nodes ui and vi. Return the center of the given star graph.

**Example 1:**

![https://assets.leetcode.com/uploads/2021/02/24/star_graph.png:w](https://assets.leetcode.com/uploads/2021/02/24/star_graph.png)

    Input: edges = [[1,2],[2,3],[4,2]]
    Output: 2
    Explanation: As shown in the figure above, node 2 is connected to every other node, so 2 is the center.

**Example 2:**

    Input: edges = [[1,2],[5,1],[1,3],[1,4]]
    Output: 1

**Constraints:**

- 3 <= n <= 100000
- edges.length == n - 1
- edges[i].length == 2
- 1 <= ui, vi <= n
- ui != vi
- The given edges represent a valid star graph.

## 题目大意

有一个无向的 星型 图，由 n 个编号从 1 到 n 的节点组成。星型图有一个 中心 节点，并且恰有 n - 1 条边将中心节点与其他每个节点连接起来。

给你一个二维整数数组 edges ，其中 edges[i] = [ui, vi] 表示在节点 ui 和 vi 之间存在一条边。请你找出并返回 edges 所表示星型图的中心节点。

## 解题思路

- 求出edges中前两个元素的共同值，即是中心节点

## 代码

```go
package leetcode

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}
```
