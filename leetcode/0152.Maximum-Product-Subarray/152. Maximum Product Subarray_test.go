package leetcode

import (
	"fmt"
	"testing"
)

type question152 struct {
	para152
	ans152
}

// para 是参数
// one 代表第一个参数
type para152 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans152 struct {
	one int
}

func Test_Problem152(t *testing.T) {

	qs := []question152{

		question152{
			para152{[]int{-2}},
			ans152{-2},
		},

		question152{
			para152{[]int{3, -1, 4}},
			ans152{4},
		},

		question152{
			para152{[]int{0}},
			ans152{0},
		},

		question152{
			para152{[]int{2, 3, -2, 4}},
			ans152{6},
		},

		question152{
			para152{[]int{-2, 0, -1}},
			ans152{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 152------------------------\n")

	for _, q := range qs {
		_, p := q.ans152, q.para152
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxProduct(p.one))
	}
	fmt.Printf("\n\n\n")
}
