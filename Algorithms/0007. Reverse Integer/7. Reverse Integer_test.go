package leetcode

import (
	"fmt"
	"testing"
)

type question7 struct {
	para7
	ans7
}

// para 是参数
// one 代表第一个参数
type para7 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans7 struct {
	one int
}

func Test_Problem7(t *testing.T) {

	qs := []question7{

		question7{
			para7{321},
			ans7{123},
		},

		question7{
			para7{-123},
			ans7{-321},
		},

		question7{
			para7{120},
			ans7{21},
		},

		question7{
			para7{1534236469},
			ans7{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 7------------------------\n")

	for _, q := range qs {
		_, p := q.ans7, q.para7
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, reverse7(p.one))
	}
	fmt.Printf("\n\n\n")
}
