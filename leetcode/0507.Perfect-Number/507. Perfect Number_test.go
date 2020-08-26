package leetcode

import (
	"fmt"
	"testing"
)

type question507 struct {
	para507
	ans507
}

// para 是参数
// one 代表第一个参数
type para507 struct {
	num int
}

// ans 是答案
// one 代表第一个答案
type ans507 struct {
	one bool
}

func Test_Problem507(t *testing.T) {

	qs := []question507{

		{
			para507{28},
			ans507{true},
		},

		{
			para507{496},
			ans507{true},
		},

		{
			para507{500},
			ans507{false},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 507------------------------\n")

	for _, q := range qs {
		_, p := q.ans507, q.para507
		fmt.Printf("【input】:%v       【output】:%v\n", p, checkPerfectNumber(p.num))
	}
	fmt.Printf("\n\n\n")
}
