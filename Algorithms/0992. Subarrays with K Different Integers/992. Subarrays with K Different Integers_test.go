package leetcode

import (
	"fmt"
	"testing"
)

type question992 struct {
	para992
	ans992
}

// para 是参数
// one 代表第一个参数
type para992 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans992 struct {
	one int
}

func Test_Problem992(t *testing.T) {

	qs := []question992{

		question992{
			para992{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 1},
			ans992{36},
		},

		question992{
			para992{[]int{2, 1, 1, 1, 2}, 1},
			ans992{8},
		},

		question992{
			para992{[]int{1, 2}, 1},
			ans992{2},
		},

		question992{
			para992{[]int{1, 2, 1, 2, 3}, 2},
			ans992{7},
		},

		question992{
			para992{[]int{1, 2, 1, 3, 4}, 3},
			ans992{3},
		},

		question992{
			para992{[]int{1}, 5},
			ans992{1},
		},

		question992{
			para992{[]int{}, 10},
			ans992{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 992------------------------\n")

	for _, q := range qs {
		_, p := q.ans992, q.para992
		fmt.Printf("【input】:%v       【output】:%v\n", p, subarraysWithKDistinct(p.one, p.k))
	}
	fmt.Printf("\n\n\n")
}
