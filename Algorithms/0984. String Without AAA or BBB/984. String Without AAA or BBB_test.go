package leetcode

import (
	"fmt"
	"testing"
)

type question984 struct {
	para984
	ans984
}

// para 是参数
// one 代表第一个参数
type para984 struct {
	a int
	b int
}

// ans 是答案
// one 代表第一个答案
type ans984 struct {
	one string
}

func Test_Problem984(t *testing.T) {

	qs := []question984{

		question984{
			para984{1, 2},
			ans984{"abb"},
		},

		question984{
			para984{4, 1},
			ans984{"aabaa"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 984------------------------\n")

	for _, q := range qs {
		_, p := q.ans984, q.para984
		fmt.Printf("【input】:%v       【output】:%v\n", p, strWithout3a3b(p.a, p.b))
	}
	fmt.Printf("\n\n\n")
}
