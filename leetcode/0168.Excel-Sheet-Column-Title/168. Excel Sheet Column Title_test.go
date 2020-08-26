package leetcode

import (
	"fmt"
	"testing"
)

type question168 struct {
	para168
	ans168
}

// para 是参数
// one 代表第一个参数
type para168 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans168 struct {
	one string
}

func Test_Problem168(t *testing.T) {

	qs := []question168{

		{
			para168{1},
			ans168{"A"},
		},

		{
			para168{28},
			ans168{"AB"},
		},

		{
			para168{701},
			ans168{"ZY"},
		},

		{
			para168{10011},
			ans168{"NUA"},
		},

		{
			para168{999},
			ans168{"ALK"},
		},

		{
			para168{681},
			ans168{"ZE"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 168------------------------\n")

	for _, q := range qs {
		_, p := q.ans168, q.para168
		fmt.Printf("【input】:%v    【output】:%v\n", p, convertToTitle(p.n))
	}
	fmt.Printf("\n\n\n")
}
