# [870. Advantage Shuffle](https://leetcode.com/problems/advantage-shuffle/)

## 题目

Given two arrays `A` and `B` of equal size, the *advantage of `A` with respect to `B`* is the number of indices `i` for which `A[i] > B[i]`.

Return **any** permutation of `A` that maximizes its advantage with respect to `B`.

**Example 1:**

```
Input:A = [2,7,11,15], B = [1,10,4,11]
Output:[2,11,7,15]
```

**Example 2:**

```
Input:A = [12,24,8,32], B = [13,25,32,11]
Output:[24,32,8,12]
```

**Note:**

1. `1 <= A.length = B.length <= 10000`
2. `0 <= A[i] <= 10^9`
3. `0 <= B[i] <= 10^9`

## 题目大意

给定两个大小相等的数组 A 和 B，A 相对于 B 的优势可以用满足 A[i] > B[i] 的索引 i 的数目来描述。返回 A 的任意排列，使其相对于 B 的优势最大化。

## 解题思路

- 此题用贪心算法解题。如果 A 中最小的牌 a 能击败 B 中最小的牌 b，那么将它们配对。否则， a 将无益于我们的比分，因为它无法击败任何牌。这是贪心的策略，每次匹配都用手中最弱的牌和 B 中的最小牌 b 进行配对，这样会使 A 中剩余的牌严格的变大，最后会使得得分更多。
- 在代码实现中，将 A 数组排序，B 数组按照下标排序。因为最终输出的是相对于 B 的优势结果，所以要针对 B 的下标不变来安排 A 的排列。排好序以后按照贪心策略选择 A 中牌的顺序。

## 代码

```go
package leetcode

import "sort"

func advantageCount1(A []int, B []int) []int {
	n := len(A)
	sort.Ints(A)
	sortedB := make([]int, n)
	for i := range sortedB {
		sortedB[i] = i
	}
	sort.Slice(sortedB, func(i, j int) bool {
		return B[sortedB[i]] < B[sortedB[j]]
	})
	useless, i, res := make([]int, 0), 0, make([]int, n)
	for _, index := range sortedB {
		b := B[index]
		for i < n && A[i] <= b {
			useless = append(useless, A[i])
			i++
		}
		if i < n {
			res[index] = A[i]
			i++
		} else {
			res[index] = useless[0]
			useless = useless[1:]
		}
	}
	return res
}
```