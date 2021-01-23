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
