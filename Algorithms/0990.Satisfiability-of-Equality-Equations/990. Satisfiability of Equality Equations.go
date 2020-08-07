package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

func equationsPossible(equations []string) bool {
	if len(equations) == 0 {
		return false
	}
	uf := template.UnionFind{}
	uf.Init(26)
	for _, equ := range equations {
		if equ[1] == '=' && equ[2] == '=' {
			uf.Union(int(equ[0]-'a'), int(equ[3]-'a'))
		}
	}
	for _, equ := range equations {
		if equ[1] == '!' && equ[2] == '=' {
			if uf.Find(int(equ[0]-'a')) == uf.Find(int(equ[3]-'a')) {
				return false
			}
		}
	}
	return true
}
