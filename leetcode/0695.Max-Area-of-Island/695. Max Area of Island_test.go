package leetcode

import (
	"fmt"
	"testing"
)

type question695 struct {
	para695
	ans695
}

// para 是参数
// one 代表第一个参数
type para695 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans695 struct {
	one int
}

func Test_Problem695(t *testing.T) {

	qs := []question695{

		{
			para695{[][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			}},
			ans695{9},
		},

		{
			para695{[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			}},
			ans695{4},
		},

		{
			para695{[][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			}},
			ans695{23},
		},

		{
			para695{[][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			}},
			ans695{5},
		},

		{
			para695{[][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			}},
			ans695{24},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 695------------------------\n")

	for _, q := range qs {
		_, p := q.ans695, q.para695
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxAreaOfIsland(p.one))
	}
	fmt.Printf("\n\n\n")
}
