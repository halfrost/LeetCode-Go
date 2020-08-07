package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

func findRedundantConnection(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	uf, res := template.UnionFind{}, []int{}
	uf.Init(len(edges) + 1)
	for i := 0; i < len(edges); i++ {
		if uf.Find(edges[i][0]) != uf.Find(edges[i][1]) {
			uf.Union(edges[i][0], edges[i][1])
		} else {
			res = append(res, edges[i][0])
			res = append(res, edges[i][1])
		}
	}
	return res
}
