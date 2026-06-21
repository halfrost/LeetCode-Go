package leetcode

import (
	"fmt"
	"testing"
)

type question1254 struct {
	para1254
	ans1254
}

// para 是参数
// one 代表第一个参数
type para1254 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1254 struct {
	one int
}

func Test_Problem1254(t *testing.T) {

	qs := []question1254{

		{
			para1254{[][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			}},
			ans1254{0},
		},

		{
			para1254{[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			}},
			ans1254{0},
		},

		{
			para1254{[][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			}},
			ans1254{2},
		},

		{
			para1254{[][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			}},
			ans1254{1},
		},

		{
			para1254{[][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			}},
			ans1254{2},
		},

		{
			para1254{[][]int{}},
			ans1254{0},
		},

		{
			para1254{[][]int{{}}},
			ans1254{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1254------------------------\n")

	for _, q := range qs {
		a, p := q.ans1254, q.para1254
		got := closedIsland(p.one)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
