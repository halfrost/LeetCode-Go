package leetcode

import (
	"fmt"
	"testing"
)

type question1281 struct {
	para1281
	ans1281
}

// para 是参数
// one 代表第一个参数
type para1281 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1281 struct {
	one int
}

func Test_Problem1281(t *testing.T) {

	qs := []question1281{

		{
			para1281{234},
			ans1281{15},
		},

		{
			para1281{4421},
			ans1281{21},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1281------------------------\n")

	for _, q := range qs {
		_, p := q.ans1281, q.para1281
		fmt.Printf("【input】:%v       【output】:%v\n", p, subtractProductAndSum(p.n))
	}
	fmt.Printf("\n\n\n")
}
