package leetcode

import (
	"fmt"
	"testing"
)

type question274 struct {
	para274
	ans274
}

// para 是参数
// one 代表第一个参数
type para274 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans274 struct {
	one int
}

func Test_Problem274(t *testing.T) {

	qs := []question274{

		question274{
			para274{[]int{3, 6, 9, 1}},
			ans274{3},
		},
		question274{
			para274{[]int{1}},
			ans274{1},
		},

		question274{
			para274{[]int{}},
			ans274{0},
		},

		question274{
			para274{[]int{3, 0, 6, 1, 5}},
			ans274{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 274------------------------\n")

	for _, q := range qs {
		_, p := q.ans274, q.para274
		fmt.Printf("【input】:%v       【output】:%v\n", p, hIndex(p.one))
	}
	fmt.Printf("\n\n\n")
}
