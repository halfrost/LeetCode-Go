package leetcode

import (
	"fmt"
	"testing"
)

type question474 struct {
	para474
	ans474
}

// para 是参数
// one 代表第一个参数
type para474 struct {
	strs []string
	m    int
	n    int
}

// ans 是答案
// one 代表第一个答案
type ans474 struct {
	one int
}

func Test_Problem474(t *testing.T) {

	qs := []question474{

		question474{
			para474{[]string{"10", "0001", "111001", "1", "0"}, 5, 3},
			ans474{4},
		},

		question474{
			para474{[]string{"10", "0", "1"}, 1, 1},
			ans474{2},
		},

		question474{
			para474{[]string{}, 0, 0},
			ans474{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 474------------------------\n")

	for _, q := range qs {
		_, p := q.ans474, q.para474
		fmt.Printf("【input】:%v       【output】:%v\n", p, findMaxForm(p.strs, p.m, p.n))
	}
	fmt.Printf("\n\n\n")
}
