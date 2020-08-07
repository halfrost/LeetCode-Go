package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

func removeStones(stones [][]int) int {
	if len(stones) <= 1 {
		return 0
	}
	uf, rowMap, colMap := template.UnionFind{}, map[int]int{}, map[int]int{}
	uf.Init(len(stones))
	for i := 0; i < len(stones); i++ {
		if _, ok := rowMap[stones[i][0]]; ok {
			uf.Union(rowMap[stones[i][0]], i)
		} else {
			rowMap[stones[i][0]] = i
		}
		if _, ok := colMap[stones[i][1]]; ok {
			uf.Union(colMap[stones[i][1]], i)
		} else {
			colMap[stones[i][1]] = i
		}
	}
	return len(stones) - uf.TotalCount()
}
