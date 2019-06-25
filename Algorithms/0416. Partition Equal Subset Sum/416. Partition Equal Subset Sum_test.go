package leetcode

import (
	"fmt"
	"testing"
)

type question416 struct {
	para416
	ans416
}

// para 是参数
// one 代表第一个参数
type para416 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans416 struct {
	one bool
}

func Test_Problem416(t *testing.T) {

	qs := []question416{

		question416{
			para416{[]int{1, 5, 11, 5}},
			ans416{true},
		},

		question416{
			para416{[]int{1, 2, 3, 5}},
			ans416{false},
		},

		question416{
			para416{[]int{1, 2, 5}},
			ans416{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 416------------------------\n")

	for _, q := range qs {
		_, p := q.ans416, q.para416
		fmt.Printf("【input】:%v       【output】:%v\n", p, canPartition(p.one))
	}
	fmt.Printf("\n\n\n")
}
