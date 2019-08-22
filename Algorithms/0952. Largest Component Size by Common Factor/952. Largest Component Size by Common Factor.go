package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

// 解法一 并查集 UnionFind
func largestComponentSize(A []int) int {
	maxElement, uf, countMap, res := 0, template.UnionFind{}, map[int]int{}, 1
	for _, v := range A {
		maxElement = max(maxElement, v)
	}
	uf.Init(maxElement + 1)
	for _, v := range A {
		for k := 2; k*k <= v; k++ {
			if v%k == 0 {
				uf.Union(v, k)
				uf.Union(v, v/k)
			}
		}
	}
	for _, v := range A {
		countMap[uf.Find(v)]++
		res = max(res, countMap[uf.Find(v)])
	}
	return res
}

// 解法二 UnionFindCount
func largestComponentSize1(A []int) int {
	uf, factorMap := template.UnionFindCount{}, map[int]int{}
	uf.Init(len(A))
	for i, v := range A {
		for k := 2; k*k <= v; k++ {
			if v%k == 0 {
				if _, ok := factorMap[k]; !ok {
					factorMap[k] = i
				} else {
					uf.Union(i, factorMap[k])
				}
				if _, ok := factorMap[v/k]; !ok {
					factorMap[v/k] = i
				} else {
					uf.Union(i, factorMap[v/k])
				}
			}
		}
		if _, ok := factorMap[v]; !ok {
			factorMap[v] = i
		} else {
			uf.Union(i, factorMap[v])
		}
	}
	return uf.MaxUnionCount()
}
