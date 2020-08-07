# [834. Sum of Distances in Tree](https://leetcode.com/problems/sum-of-distances-in-tree/)


## 题目:

An undirected, connected tree with `N` nodes labelled `0...N-1` and `N-1edges` are given.

The `i`th edge connects nodes `edges[i][0]` and `edges[i][1]` together.

Return a list `ans`, where `ans[i]` is the sum of the distances between node `i`and all other nodes.

**Example 1:**

    Input: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
    Output: [8,12,6,10,10,10]
    Explanation: 
    Here is a diagram of the given tree:
      0
     / \
    1   2
       /|\
      3 4 5
    We can see that dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
    equals 1 + 1 + 2 + 2 + 2 = 8.  Hence, answer[0] = 8, and so on.

Note: `1 <= N <= 10000`

## 题目大意

给定一个无向、连通的树。树中有 N 个标记为 0...N-1 的节点以及 N-1 条边。第 i 条边连接节点 edges[i][0] 和 edges[i][1] 。返回一个表示节点 i 与其他所有节点距离之和的列表 ans。

说明: 1 <= N <= 10000



## 解题思路

- 给出 N 个节点和这些节点之间的一些边的关系。要求求出分别以 x 为根节点到所有节点路径和。
- 这一题虽说描述的是求树的路径，但是完全可以当做图来做，因为并不是二叉树，是多叉树。这一题的解题思路是先一次 DFS 求出以 0 为根节点到各个节点的路径和(不以 0 为节点也可以，可以取任意节点作为开始)。第二次 DFS 求出从 0 根节点转换到其他各个节点的路径和。由于第一次计算出来以 0 为节点的路径和是正确的，所以计算其他节点为根节点的路径和只需要转换一下就可以得到正确结果。经过 2 次 DFS 之后就可以得到所有节点以自己为根节点到所有节点的路径和了。
- 如何从以 0 为根节点到其他所有节点的路径和转换到以其他节点为根节点到所有节点的路径和呢？从 0 节点换成 x 节点，只需要在 0 到所有节点的路径和基础上增增减减就可以了。增加的是 x 节点到除去以 x 为根节点所有子树以外的节点的路径，有多少个节点就增加多少条路径。减少的是 0 到以 x 为根节点所有子树节点的路径和，包含 0 到 x 根节点，有多少节点就减少多少条路径。所以在第一次 DFS 中需要计算好每个节点以自己为根节点的子树总数和(包含自己在内)，这样在第二次 DFS 中可以直接拿来做转换。具体细节的实现见代码。

