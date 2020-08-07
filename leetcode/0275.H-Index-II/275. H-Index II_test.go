package leetcode

import (
	"fmt"
	"testing"
)

type question275 struct {
	para275
	ans275
}

// para 是参数
// one 代表第一个参数
type para275 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans275 struct {
	one int
}

func Test_Problem275(t *testing.T) {

	qs := []question275{

		question275{
			para275{[]int{3, 6, 9, 1}},
			ans275{3},
		},
		question275{
			para275{[]int{1}},
			ans275{1},
		},

		question275{
			para275{[]int{}},
			ans275{0},
		},

		question275{
			para275{[]int{3, 0, 6, 1, 5}},
			ans275{3},
		},

		question275{
			para275{[]int{0, 1, 3, 5, 6}},
			ans275{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 275------------------------\n")

	for _, q := range qs {
		_, p := q.ans275, q.para275
		fmt.Printf("【input】:%v       【output】:%v\n", p, hIndex275(p.one))
	}
	fmt.Printf("\n\n\n")
}
