package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/template"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	uf, res, sMap := template.UnionFind{}, []byte(s), map[int][]byte{}
	uf.Init(len(s))
	for _, pair := range pairs {
		uf.Union(pair[0], pair[1])
	}
	for i := 0; i < len(s); i++ {
		r := uf.Find(i)
		sMap[r] = append(sMap[r], s[i])
	}
	for _, v := range sMap {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}
	for i := 0; i < len(s); i++ {
		r := uf.Find(i)
		bytes := sMap[r]
		res[i] = bytes[0]
		sMap[r] = bytes[1:len(bytes)]
	}
	return string(res)
}
