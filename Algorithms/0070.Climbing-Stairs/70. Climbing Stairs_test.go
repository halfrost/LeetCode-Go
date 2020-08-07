package leetcode

import (
	"fmt"
	"testing"
)

type question70 struct {
	para70
	ans70
}

// para 是参数
// one 代表第一个参数
type para70 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans70 struct {
	one int
}

func Test_Problem70(t *testing.T) {

	qs := []question70{

		question70{
			para70{2},
			ans70{2},
		},

		question70{
			para70{3},
			ans70{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 70------------------------\n")

	for _, q := range qs {
		_, p := q.ans70, q.para70
		fmt.Printf("【input】:%v       【output】:%v\n", p, climbStairs(p.n))
	}
	fmt.Printf("\n\n\n")
}
