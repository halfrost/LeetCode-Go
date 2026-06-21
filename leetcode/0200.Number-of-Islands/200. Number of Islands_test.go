package leetcode

import (
	"fmt"
	"testing"
)

type question200 struct {
	para200
	ans200
}

// para 是参数
// one 代表第一个参数
type para200 struct {
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans200 struct {
	one int
}

func Test_Problem200(t *testing.T) {

	qs := []question200{

		{
			para200{[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}},
			ans200{1},
		},

		{
			para200{[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			}},
			ans200{3},
		},

		{
			para200{[][]byte{}},
			ans200{0},
		},

		{
			para200{[][]byte{{}}},
			ans200{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 200------------------------\n")

	for _, q := range qs {
		a, p := q.ans200, q.para200
		out := numIslands(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, out)
		if out != a.one {
			t.Fatalf("input %v expected %v got %v", p, a.one, out)
		}
	}
	fmt.Printf("\n\n\n")
}
