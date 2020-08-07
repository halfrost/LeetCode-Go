package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

func numSimilarGroups(A []string) int {
	uf := template.UnionFind{}
	uf.Init(len(A))
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if isSimilar(A[i], A[j]) {
				uf.Union(i, j)
			}
		}
	}
	return uf.TotalCount()
}

func isSimilar(a, b string) bool {
	var n int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			n++
			if n > 2 {
				return false
			}
		}
	}
	return true
}
