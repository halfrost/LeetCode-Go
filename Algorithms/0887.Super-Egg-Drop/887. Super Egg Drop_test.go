package leetcode

import (
	"fmt"
	"testing"
)

type question887 struct {
	para887
	ans887
}

// para 是参数
// one 代表第一个参数
type para887 struct {
	k int
	n int
}

// ans 是答案
// one 代表第一个答案
type ans887 struct {
	one int
}

func Test_Problem887(t *testing.T) {

	qs := []question887{

		question887{
			para887{1, 2},
			ans887{2},
		},

		question887{
			para887{2, 6},
			ans887{3},
		},

		question887{
			para887{3, 14},
			ans887{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 887------------------------\n")

	for _, q := range qs {
		_, p := q.ans887, q.para887
		fmt.Printf("【input】:%v       【output】:%v\n", p, superEggDrop(p.k, p.n))
	}
	fmt.Printf("\n\n\n")
}
