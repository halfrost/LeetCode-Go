package leetcode

import (
	"fmt"
	"testing"
)

type question1678 struct {
	para1678
	ans1678
}

// para 是参数
// one 代表第一个参数
type para1678 struct {
	command string
}

// ans 是答案
// one 代表第一个答案
type ans1678 struct {
	one string
}

func Test_Problem1678(t *testing.T) {

	qs := []question1678{

		{
			para1678{"G()(al)"},
			ans1678{"Goal"},
		},

		{
			para1678{"G()()()()(al)"},
			ans1678{"Gooooal"},
		},

		{
			para1678{"(al)G(al)()()G"},
			ans1678{"alGalooG"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1678------------------------\n")

	for _, q := range qs {
		_, p := q.ans1678, q.para1678
		fmt.Printf("【input】:%v       【output】:%v\n", p, interpret(p.command))
	}
	fmt.Printf("\n\n\n")
}
