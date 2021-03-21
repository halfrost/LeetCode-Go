# [1319. Number of Operations to Make Network Connected](https://leetcode.com/problems/number-of-operations-to-make-network-connected/)


## 题目

There are `n` computers numbered from `0` to `n-1` connected by ethernet cables `connections` forming a network where `connections[i] = [a, b]` represents a connection between computers `a` and `b`. Any computer can reach any other computer directly or indirectly through the network.

Given an initial computer network `connections`. You can extract certain cables between two directly connected computers, and place them between any pair of disconnected computers to make them directly connected. Return the *minimum number of times* you need to do this in order to make all the computers connected. If it's not possible, return -1.

**Example 1:**

![https://assets.leetcode.com/uploads/2020/01/02/sample_1_1677.png](https://assets.leetcode.com/uploads/2020/01/02/sample_1_1677.png)

```
Input: n = 4, connections = [[0,1],[0,2],[1,2]]
Output: 1
Explanation: Remove cable between computer 1 and 2 and place between computers 1 and 3.
```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/01/02/sample_2_1677.png](https://assets.leetcode.com/uploads/2020/01/02/sample_2_1677.png)

```
Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
Output: 2
```

**Example 3:**

```
Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
Output: -1
Explanation: There are not enough cables.
```

**Example 4:**

```
Input: n = 5, connections = [[0,1],[0,2],[3,4],[2,3]]
Output: 0
```

**Constraints:**

- `1 <= n <= 10^5`
- `1 <= connections.length <= min(n*(n-1)/2, 10^5)`
- `connections[i].length == 2`
- `0 <= connections[i][0], connections[i][1] < n`
- `connections[i][0] != connections[i][1]`
- There are no repeated connections.
- No two computers are connected by more than one cable.

## 题目大意

用以太网线缆将 n 台计算机连接成一个网络，计算机的编号从 0 到 n-1。线缆用 connections 表示，其中 connections[i] = [a, b] 连接了计算机 a 和 b。网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。

## 解题思路

- 很明显这题的解题思路是并查集。先将每个 connections 构建出并查集。构建中需要累加冗余的连接。例如 2 个节点已经连通，再连接这个集合中的任意 2 个节点就算冗余连接。冗余连接的线都可以移动，去连接还没有连通的节点。计算出冗余连接数，再根据并查集的集合总数，即可得出答案。
- 这一题答案有 3 种可能。第一种，所有点都在一个集合内，即全部连通，这时输出 0 。第二种，冗余的连接不够串起所有点，这时输出 -1 。第三种情况是可以连通的情况。 m 个集合需要连通，最少需要 m - 1 条线。如果冗余连接数大于 m - 1，则输出 m - 1 即可。

## 代码

```go
package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

func makeConnected(n int, connections [][]int) int {
	if n-1 > len(connections) {
		return -1
	}
	uf, redundance := template.UnionFind{}, 0
	uf.Init(n)
	for _, connection := range connections {
		if uf.Find(connection[0]) == uf.Find(connection[1]) {
			redundance++
		} else {
			uf.Union(connection[0], connection[1])
		}
	}
	if uf.TotalCount() == 1 || redundance < uf.TotalCount()-1 {
		return 0
	}
	return uf.TotalCount() - 1
}
```