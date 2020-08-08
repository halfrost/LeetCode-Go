# [851. Loud and Rich](https://leetcode.com/problems/loud-and-rich/)



## 题目

In a group of N people (labelled `0, 1, 2, ..., N-1`), each person has different amounts of money, and different levels of quietness.

For convenience, we'll call the person with label `x`, simply "person `x`".

We'll say that `richer[i] = [x, y]` if person `x` definitely has more money than person `y`. Note that `richer` may only be a subset of valid observations.

Also, we'll say `quiet = q` if person x has quietness `q`.

Now, return `answer`, where `answer = y` if `y` is the least quiet person (that is, the person `y` with the smallest value of `quiet[y]`), among all people who definitely have equal to or more money than person `x`.

**Example 1**:

```
Input: richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,0]
Output: [5,5,2,5,4,5,6,7]
Explanation: 
answer[0] = 5.
Person 5 has more money than 3, which has more money than 1, which has more money than 0.
The only person who is quieter (has lower quiet[x]) is person 7, but
it isn't clear if they have more money than person 0.

answer[7] = 7.
Among all people that definitely have equal to or more money than person 7
(which could be persons 3, 4, 5, 6, or 7), the person who is the quietest (has lower quiet[x])
is person 7.

The other answers can be filled out with similar reasoning.
```

**Note**:

1. `1 <= quiet.length = N <= 500`
2. `0 <= quiet[i] < N`, all `quiet[i]` are different.
3. `0 <= richer.length <= N * (N-1) / 2`
4. `0 <= richer[i][j] < N`
5. `richer[i][0] != richer[i][1]`
6. `richer[i]`'s are all different.
7. The observations in `richer` are all logically consistent.

## 题目大意

在一组 N 个人（编号为 0, 1, 2, ..., N-1）中，每个人都有不同数目的钱，以及不同程度的安静（quietness）。为了方便起见，我们将编号为 x 的人简称为 "person x "。如果能够肯定 person x 比 person y 更有钱的话，我们会说 richer[i] = [x, y] 。注意 richer 可能只是有效观察的一个子集。另外，如果 person x 的安静程度为 q ，我们会说 quiet[x] = q 。现在，返回答案 answer ，其中 answer[x] = y 的前提是，在所有拥有的钱不少于 person x 的人中，person y 是最安静的人（也就是安静值 quiet[y] 最小的人）。

提示：

- 1 <= quiet.length = N <= 500
- 0 <= quiet[i] < N，所有 quiet[i] 都不相同。
- 0 <= richer.length <= N * (N-1) / 2
- 0 <= richer[i][j] < N
- richer[i][0] != richer[i][1]
- richer[i] 都是不同的。
- 对 richer 的观察在逻辑上是一致的。


## 解题思路

- 给出 2 个数组，richer 和 quiet，要求输出 answer，其中 answer = y 的前提是，在所有拥有的钱不少于 x 的人中，y 是最安静的人（也就是安静值 quiet[y] 最小的人）
- 由题意可知，`richer` 构成了一个有向无环图，首先使用字典建立图的关系，找到比当前下标编号富有的所有的人。然后使用广度优先层次遍历，不断的使用富有的人，但是安静值更小的人更新子节点即可。
- 这一题还可以用拓扑排序来解答。将 `richer` 中描述的关系看做边，如果 `x > y`，则 `x` 指向 `y`。将 `quiet` 看成权值。用一个数组记录答案，初始时 `ans[i] = i`。然后对原图做拓扑排序，对于每一条边，如果发现 `quiet[ans[v]] > quiet[ans[u]]`，则 `ans[v]` 的答案为 `ans[u]`。时间复杂度即为拓扑排序的时间复杂度为 `O(m+n)`。空间复杂度需要 `O(m)` 的数组建图，需要 `O(n)` 的数组记录入度以及存储队列，所以空间复杂度为 `O(m+n)`。

## 代码

```go
func loudAndRich(richer [][]int, quiet []int) []int {
	edges := make([][]int, len(quiet))
	for i := range edges {
		edges[i] = []int{}
	}
	indegrees := make([]int, len(quiet))
	for _, edge := range richer {
		n1, n2 := edge[0], edge[1]
		edges[n1] = append(edges[n1], n2)
		indegrees[n2]++
	}
	res := make([]int, len(quiet))
	for i := range res {
		res[i] = i
	}
	queue := []int{}
	for i, v := range indegrees {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		nexts := []int{}
		for _, n1 := range queue {
			for _, n2 := range edges[n1] {
				indegrees[n2]--
				if quiet[res[n2]] > quiet[res[n1]] {
					res[n2] = res[n1]
				}
				if indegrees[n2] == 0 {
					nexts = append(nexts, n2)
				}
			}
		}
		queue = nexts
	}
	return res
}
```