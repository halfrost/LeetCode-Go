package leetcode

import (
	"fmt"
	"testing"
)

type question483 struct {
	para483
	ans483
}

// para 是参数
// one 代表第一个参数
type para483 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans483 struct {
	one string
}

func Test_Problem483(t *testing.T) {

	qs := []question483{

		question483{
			para483{"13"},
			ans483{"3"},
		},

		question483{
			para483{"4681"},
			ans483{"8"},
		},

		question483{
			para483{"1000000000000000000"},
			ans483{"999999999999999999"},
		},

		question483{
			para483{"727004545306745403"},
			ans483{"727004545306745402"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 483------------------------\n")

	for _, q := range qs {
		_, p := q.ans483, q.para483
		fmt.Printf("【input】:%v       【output】:%v\n", p, smallestGoodBase(p.one))
	}
	fmt.Printf("\n\n\n")
}
