package leetcode

import (
	"fmt"
	"testing"
)

type question67 struct {
	para67
	ans67
}

// para 是参数
// one 代表第一个参数
type para67 struct {
	a string
	b string
}

// ans 是答案
// one 代表第一个答案
type ans67 struct {
	one string
}

func Test_Problem67(t *testing.T) {

	qs := []question67{

		{
			para67{"11", "1"},
			ans67{"100"},
		},

		{
			para67{"1010", "1011"},
			ans67{"10101"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 67------------------------\n")

	for _, q := range qs {
		_, p := q.ans67, q.para67
		fmt.Printf("【input】:%v       【output】:%v\n", p, addBinary(p.a, p.b))
	}
	fmt.Printf("\n\n\n")
}
