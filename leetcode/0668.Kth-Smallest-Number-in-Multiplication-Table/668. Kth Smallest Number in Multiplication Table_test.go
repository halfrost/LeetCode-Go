package leetcode

import (
	"fmt"
	"testing"
)

type question668 struct {
	para668
	ans668
}

// para 是参数
// one 代表第一个参数
type para668 struct {
	m int
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans668 struct {
	one int
}

func Test_Problem668(t *testing.T) {

	qs := []question668{

		question668{
			para668{3, 3, 5},
			ans668{3},
		},

		question668{
			para668{2, 3, 6},
			ans668{6},
		},

		question668{
			para668{1, 3, 2},
			ans668{2},
		},

		question668{
			para668{42, 34, 401},
			ans668{126},
		},

		question668{
			para668{7341, 13535, 12330027},
			ans668{2673783},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 668------------------------\n")

	for _, q := range qs {
		_, p := q.ans668, q.para668
		fmt.Printf("【input】:%v       【output】:%v\n", p, findKthNumber(p.m, p.n, p.k))
	}
	fmt.Printf("\n\n\n")
}
