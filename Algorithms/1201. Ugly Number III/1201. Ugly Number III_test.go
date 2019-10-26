package leetcode

import (
	"fmt"
	"testing"
)

type question1201 struct {
	para1201
	ans1201
}

// para 是参数
// one 代表第一个参数
type para1201 struct {
	n int
	a int
	b int
	c int
}

// ans 是答案
// one 代表第一个答案
type ans1201 struct {
	one int
}

func Test_Problem1201(t *testing.T) {

	qs := []question1201{

		question1201{
			para1201{3, 2, 3, 5},
			ans1201{4},
		},

		question1201{
			para1201{4, 2, 3, 4},
			ans1201{6},
		},

		question1201{
			para1201{5, 2, 11, 13},
			ans1201{10},
		},

		question1201{
			para1201{1000000000, 2, 217983653, 336916467},
			ans1201{1999999984},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1201------------------------\n")

	for _, q := range qs {
		_, p := q.ans1201, q.para1201
		fmt.Printf("【input】:%v       【output】:%v\n", p, nthUglyNumber(p.n, p.a, p.b, p.c))
	}
	fmt.Printf("\n\n\n")
}
