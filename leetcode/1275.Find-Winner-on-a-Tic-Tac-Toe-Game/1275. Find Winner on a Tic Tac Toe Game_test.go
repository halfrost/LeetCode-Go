package leetcode

import (
	"fmt"
	"testing"
)

type question1275 struct {
	para1275
	ans1275
}

// para 是参数
// one 代表第一个参数
type para1275 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1275 struct {
	one string
}

func Test_Problem1275(t *testing.T) {

	qs := []question1275{

		{
			para1275{[][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}},
			ans1275{"A"},
		},

		{
			para1275{[][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}},
			ans1275{"B"},
		},

		{
			para1275{[][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}},
			ans1275{"Draw"},
		},

		{
			para1275{[][]int{{0, 0}, {1, 1}}},
			ans1275{"Pending"},
		},

		// A wins by row 0
		{
			para1275{[][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}}},
			ans1275{"A"},
		},

		// B wins by row 1
		{
			para1275{[][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {2, 2}, {1, 2}}},
			ans1275{"B"},
		},

		// A wins by column 0
		{
			para1275{[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}},
			ans1275{"A"},
		},

		// B wins by column 1
		{
			para1275{[][]int{{0, 0}, {0, 1}, {0, 2}, {1, 1}, {2, 2}, {2, 1}}},
			ans1275{"B"},
		},

		// B wins by main diagonal
		{
			para1275{[][]int{{0, 1}, {0, 0}, {0, 2}, {1, 1}, {1, 0}, {2, 2}}},
			ans1275{"B"},
		},

		// A wins by anti diagonal
		{
			para1275{[][]int{{0, 2}, {0, 0}, {1, 1}, {0, 1}, {2, 0}}},
			ans1275{"A"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1275------------------------\n")

	for _, q := range qs {
		a, p := q.ans1275, q.para1275
		got := tictactoe(p.one)
		if got != a.one {
			t.Fatalf("input %v: got %v, want %v", p.one, got, a.one)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
