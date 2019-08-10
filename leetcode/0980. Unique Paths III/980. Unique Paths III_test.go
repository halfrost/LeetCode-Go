package leetcode

import (
	"fmt"
	"testing"
)

type question980 struct {
	para980
	ans980
}

// para 是参数
// one 代表第一个参数
type para980 struct {
	grid [][]int
}

// ans 是答案
// one 代表第一个答案
type ans980 struct {
	one int
}

func Test_Problem980(t *testing.T) {

	qs := []question980{

		question980{
			para980{[][]int{
				[]int{1, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 2, -1},
			}},
			ans980{2},
		},

		question980{
			para980{[][]int{
				[]int{1, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 2},
			}},
			ans980{4},
		},

		question980{
			para980{[][]int{
				[]int{1, 0},
				[]int{0, 2},
			}},
			ans980{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 980------------------------\n")

	for _, q := range qs {
		_, p := q.ans980, q.para980
		fmt.Printf("【input】:%v       【output】:%v\n\n\n", p, uniquePathsIII(p.grid))
	}
	fmt.Printf("\n\n\n")
}
