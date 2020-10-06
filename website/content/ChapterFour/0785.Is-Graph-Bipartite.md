# [785. Is Graph Bipartite?](https://leetcode.com/problems/is-graph-bipartite/)


## 题目

Given an undirected `graph`, return `true` if and only if it is bipartite.

Recall that a graph is *bipartite* if we can split it's set of nodes into two independent subsets A and B such that every edge in the graph has one node in A and another node in B.

The graph is given in the following form: `graph[i]` is a list of indexes `j` for which the edge between nodes `i` and `j` exists. Each node is an integer between `0` and `graph.length - 1`. There are no self edges or parallel edges: `graph[i]` does not contain `i`, and it doesn't contain any element twice.


	Example 1:Input: [[1,3], [0,2], [1,3], [0,2]]
	Output: true
	Explanation: 
	The graph looks like this:
	0----1
	|    |
	|    |
	3----2
	We can divide the vertices into two groups: {0, 2} and {1, 3}.


	Example 2:Input: [[1,2,3], [0,2], [0,1,3], [0,2]]
	Output: false
	Explanation: 
	The graph looks like this:
	0----1
	| \  |
	|  \ |
	3----2
	We cannot find a way to divide the set of nodes into two independent subsets.


**Note:**

- `graph` will have length in range `[1, 100]`.
- `graph[i]` will contain integers in range `[0, graph.length - 1]`.
- `graph[i]` will not contain `i` or duplicate values.
- The graph is undirected: if any element `j` is in `graph[i]`, then `i` will be in `graph[j]`.

## 题目大意

给定一个无向图 graph，当这个图为二分图时返回 true。

graph 将会以邻接表方式给出，graph[i] 表示图中与节点i相连的所有节点。每个节点都是一个在 0 到 graph.length-1 之间的整数。这图中没有自环和平行边： graph[i] 中不存在 i，并且 graph[i] 中没有重复的值。

注意:

- graph 的长度范围为 [1, 100]。
- graph[i] 中的元素的范围为 [0, graph.length - 1]。
- graph[i] 不会包含 i 或者有重复的值。
- 图是无向的: 如果 j 在 graph[i] 里边, 那么 i 也会在 graph[j] 里边。

## 解题思路

- 判断一个无向图是否是二分图。二分图的定义：如果我们能将一个图的节点集合分割成两个独立的子集 A 和 B，并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，我们就将这个图称为二分图。
- 这一题可以用 BFS、DFS、并查集来解答。这里是 DFS 实现。任选一个节点开始，把它染成红色，然后对整个图 DFS 遍历，把与它相连的节点并且未被染色的，都染成绿色。颜色不同的节点代表不同的集合。这时候还可能遇到第 2 种情况，与它相连的节点已经有颜色了，并且这个颜色和前一个节点的颜色相同，这就说明了该无向图不是二分图。可以直接 return false。如此遍历到所有节点都染色了，如果能染色成功，说明该无向图是二分图，返回 true。

## 代码

```go
package leetcode

// DFS 染色，1 是红色，0 是绿色，-1 是未染色
func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	for i := range colors {
		colors[i] = -1
	}
	for i := range graph {
		if !dfs(i, graph, colors, -1) {
			return false
		}
	}
	return true
}

func dfs(n int, graph [][]int, colors []int, parentCol int) bool {
	if colors[n] == -1 {
		if parentCol == 1 {
			colors[n] = 0
		} else {
			colors[n] = 1
		}
	} else if colors[n] == parentCol {
		return false
	} else if colors[n] != parentCol {
		return true
	}
	for _, c := range graph[n] {
		if !dfs(c, graph, colors, colors[n]) {
			return false
		}
	}
	return true
}
```