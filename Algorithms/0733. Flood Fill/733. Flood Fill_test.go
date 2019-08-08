package leetcode

import (
	"fmt"
	"testing"
)

type question733 struct {
	para733
	ans733
}

// para 是参数
// one 代表第一个参数
type para733 struct {
	one [][]int
	sr  int
	sc  int
	c   int
}

// ans 是答案
// one 代表第一个答案
type ans733 struct {
	one [][]int
}

func Test_Problem733(t *testing.T) {

	qs := []question733{

		question733{
			para733{[][]int{
				[]int{1, 1, 1},
				[]int{1, 1, 0},
				[]int{1, 0, 1},
			}, 1, 1, 2},
			ans733{[][]int{
				[]int{2, 2, 2},
				[]int{2, 2, 0},
				[]int{2, 0, 1},
			}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 733------------------------\n")

	for _, q := range qs {
		_, p := q.ans733, q.para733
		fmt.Printf("【input】:%v       【output】:%v\n", p, floodFill(p.one, p.sr, p.sc, p.c))
	}
	fmt.Printf("\n\n\n")
}
