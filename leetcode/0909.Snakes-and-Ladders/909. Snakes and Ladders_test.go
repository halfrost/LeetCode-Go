package leetcode

import (
	"fmt"
	"testing"
)

type question909 struct {
	para909
	ans909
}

// para 是参数
// one 代表第一个参数
type para909 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans909 struct {
	one int
}

func Test_Problem909(t *testing.T) {
	qs := []question909{
		{
			para909{[][]int{
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 35, -1, -1, 13, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 15, -1, -1, -1, -1},
			}},
			ans909{4},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 909------------------------\n")
	for _, q := range qs {
		_, p := q.ans909, q.para909
		fmt.Printf("【input】:%v       【output】:%v\n", p, snakesAndLadders(p.one))
	}
	fmt.Printf("\n\n\n")
}
