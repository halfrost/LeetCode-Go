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
		{
			// 无法到达终点，返回 -1
			para909{[][]int{
				{1, 1, -1},
				{1, 1, 1},
				{-1, 1, 1},
			}},
			ans909{-1},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 909------------------------\n")
	for _, q := range qs {
		a, p := q.ans909, q.para909
		got := snakesAndLadders(p.one)
		if got != a.one {
			t.Fatalf("input %v: got %d, want %d", p.one, got, a.one)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
