package leetcode

import (
	"fmt"
	"testing"
)

type question838 struct {
	para838
	ans838
}

// para 是参数
// one 代表第一个参数
type para838 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans838 struct {
	one string
}

func Test_Problem838(t *testing.T) {

	qs := []question838{

		question838{
			para838{".L.R...LR..L.."},
			ans838{"LL.RR.LLRRLL.."},
		},

		question838{
			para838{"RR.L"},
			ans838{"RR.L"},
		},

		question838{
			para838{".L.R."},
			ans838{"LL.RR"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 838------------------------\n")

	for _, q := range qs {
		_, p := q.ans838, q.para838
		fmt.Printf("【input】:%v       【output】:%v\n", p, pushDominoes(p.one))
	}
	fmt.Printf("\n\n\n")
}
