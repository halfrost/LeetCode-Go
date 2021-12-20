package leetcode

import (
	"fmt"
	"testing"
)

type question997 struct {
	para997
	ans997
}

// para 是参数
type para997 struct {
	n     int
	trust [][]int
}

// ans 是答案
type ans997 struct {
	ans int
}

func Test_Problem997(t *testing.T) {

	qs := []question997{

		{
			para997{2, [][]int{{1, 2}}},
			ans997{2},
		},

		{
			para997{3, [][]int{{1, 3}, {2, 3}}},
			ans997{3},
		},

		{
			para997{3, [][]int{{1, 3}, {2, 3}, {3, 1}}},
			ans997{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 997------------------------\n")

	for _, q := range qs {
		_, p := q.ans997, q.para997
		fmt.Printf("【input】:%v    【output】:%v\n", p, findJudge(p.n, p.trust))
	}
	fmt.Printf("\n\n\n")
}
