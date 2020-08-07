package leetcode

import (
	"fmt"
	"testing"
)

type question461 struct {
	para461
	ans461
}

// para 是参数
// one 代表第一个参数
type para461 struct {
	x int
	y int
}

// ans 是答案
// one 代表第一个答案
type ans461 struct {
	one int
}

func Test_Problem461(t *testing.T) {

	qs := []question461{

		question461{
			para461{1, 4},
			ans461{2},
		},

		question461{
			para461{1, 1},
			ans461{0},
		},

		question461{
			para461{1, 3},
			ans461{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 461------------------------\n")

	for _, q := range qs {
		_, p := q.ans461, q.para461
		fmt.Printf("【input】:%v       【output】:%v\n", p, hammingDistance(p.x, p.y))
	}
	fmt.Printf("\n\n\n")
}
