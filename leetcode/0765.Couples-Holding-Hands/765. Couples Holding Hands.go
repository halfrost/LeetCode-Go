package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

func minSwapsCouples(row []int) int {
	if len(row)&1 == 1 {
		return 0
	}
	uf := template.UnionFind{}
	uf.Init(len(row))
	for i := 0; i < len(row)-1; i = i + 2 {
		uf.Union(i, i+1)
	}
	for i := 0; i < len(row)-1; i = i + 2 {
		if uf.Find(row[i]) != uf.Find(row[i+1]) {
			uf.Union(row[i], row[i+1])
		}
	}
	return len(row)/2 - uf.TotalCount()
}
