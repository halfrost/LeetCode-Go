package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

func maxNumEdgesToRemove(n int, edges [][]int) int {
	alice, bob, res := template.UnionFind{}, template.UnionFind{}, len(edges)
	alice.Init(n)
	bob.Init(n)
	for _, e := range edges {
		x, y := e[1]-1, e[2]-1
		if e[0] == 3 && (!(alice.Find(x) == alice.Find(y)) || !(bob.Find(x) == bob.Find(y))) {
			alice.Union(x, y)
			bob.Union(x, y)
			res--
		}
	}
	ufs := [2]*template.UnionFind{&alice, &bob}
	for _, e := range edges {
		if tp := e[0]; tp < 3 && !(ufs[tp-1].Find(e[1]-1) == ufs[tp-1].Find(e[2]-1)) {
			ufs[tp-1].Union(e[1]-1, e[2]-1)
			res--
		}
	}
	if alice.TotalCount() > 1 || bob.TotalCount() > 1 {
		return -1
	}
	return res
}
