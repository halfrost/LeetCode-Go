# [802. Find Eventual Safe States](https://leetcode.com/problems/find-eventual-safe-states/)



## 题目

In a directed graph, we start at some node and every turn, walk along a directed edge of the graph. If we reach a node that is terminal (that is, it has no outgoing directed edges), we stop.

Now, say our starting node is *eventually safe* if and only if we must eventually walk to a terminal node. More specifically, there exists a natural number `K` so that for any choice of where to walk, we must have stopped at a terminal node in less than `K` steps.

Which nodes are eventually safe? Return them as an array in sorted order.

The directed graph has `N` nodes with labels `0, 1, ..., N-1`, where `N` is the length of `graph`. The graph is given in the following form: `graph[i]` is a list of labels `j` such that `(i, j)` is a directed edge of the graph.

```
Example:
Input: graph = [[1,2],[2,3],[5],[0],[5],[],[]]
Output: [2,4,5,6]
Here is a diagram of the above graph.
```

![https://s3-lc-upload.s3.amazonaws.com/uploads/2018/03/17/picture1.png](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/03/17/picture1.png)

**Note**:

- `graph` will have length at most `10000`.
- The number of edges in the graph will not exceed `32000`.
- Each `graph[i]` will be a sorted list of different integers, chosen within the range `[0, graph.length - 1]`.

## 题目大意

在有向图中, 我们从某个节点和每个转向处开始, 沿着图的有向边走。 如果我们到达的节点是终点 (即它没有连出的有向边), 我们停止。现在, 如果我们最后能走到终点，那么我们的起始节点是最终安全的。 更具体地说, 存在一个自然数 K,  无论选择从哪里开始行走, 我们走了不到 K 步后必能停止在一个终点。哪些节点最终是安全的？ 结果返回一个有序的数组。

提示：

- graph 节点数不超过 10000.
- 图的边数不会超过 32000.
- 每个 graph[i] 被排序为不同的整数列表， 在区间 [0, graph.length - 1] 中选取。


## 解题思路

- 给出一个有向图，要求找出所有“安全”节点。“安全”节点的定义是：存在一个自然数 K, 无论选择从哪里开始行走, 我们走了不到 K 步后必能停止在一个终点。
- 这一题可以用拓扑排序，也可以用 DFS 染色来解答。这里用 DFS 来解答。对于每个节点，我们有 3 种染色的方法：白色 0 号节点表示该节点还没有被访问过；灰色 1 号节点表示该节点在栈中（这一轮搜索中被访问过）或者在环中；黑色 2 号节点表示该节点的所有相连的节点都被访问过，且该节点不在环中。当我们第一次访问一个节点时，我们把它从白色变成灰色，并继续搜索与它相连的节点。如果在搜索过程中我们遇到一个灰色的节点，那么说明找到了一个环，此时退出搜索，所有的灰色节点保持不变（即从任意一个灰色节点开始，都能走到环中），如果搜索过程中，我们没有遇到灰色的节点，那么在回溯到当前节点时，我们把它从灰色变成黑色，即表示它是一个安全的节点。

## 代码

```go
func eventualSafeNodes(graph [][]int) []int {
	res, color := []int{}, make([]int, len(graph))
	for i := range graph {
		if dfsEventualSafeNodes(graph, i, color) {
			res = append(res, i)
		}
	}
	return res
}

// colors: WHITE 0, GRAY 1, BLACK 2;
func dfsEventualSafeNodes(graph [][]int, idx int, color []int) bool {
	if color[idx] > 0 {
		return color[idx] == 2
	}
	color[idx] = 1
	for i := range graph[idx] {
		if !dfsEventualSafeNodes(graph, graph[idx][i], color) {
			return false
		}
	}
	color[idx] = 2
	return true
}
```