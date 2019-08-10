package leetcode

import (
	"fmt"
	"testing"
)

type question907 struct {
	para907
	ans907
}

// para 是参数
// one 代表第一个参数
type para907 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans907 struct {
	one int
}

func Test_Problem907(t *testing.T) {

	qs := []question907{

		question907{
			para907{[]int{3, 1, 2, 4}},
			ans907{17},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 907------------------------\n")

	for _, q := range qs {
		_, p := q.ans907, q.para907
		fmt.Printf("【input】:%v       【output】:%v\n", p, sumSubarrayMins(p.one))
	}
	fmt.Printf("\n\n\n")
}
