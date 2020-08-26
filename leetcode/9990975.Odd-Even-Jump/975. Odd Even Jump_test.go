package leetcode

import (
	"fmt"
	"testing"
)

type question975 struct {
	para975
	ans975
}

// para 是参数
// one 代表第一个参数
type para975 struct {
	A []int
}

// ans 是答案
// one 代表第一个答案
type ans975 struct {
	one int
}

func Test_Problem975(t *testing.T) {

	qs := []question975{

		{
			para975{[]int{10, 13, 12, 14, 15}},
			ans975{2},
		},

		{
			para975{[]int{2, 3, 1, 1, 4}},
			ans975{3},
		},

		{
			para975{[]int{5, 1, 3, 4, 2}},
			ans975{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 975------------------------\n")

	for _, q := range qs {
		_, p := q.ans975, q.para975
		fmt.Printf("【input】:%v       【output】:%v\n", p, oddEvenJumps(p.A))
	}
	fmt.Printf("\n\n\n")
}
