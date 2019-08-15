package leetcode

import (
	"fmt"
	"testing"
)

type question372 struct {
	para372
	ans372
}

// para 是参数
// one 代表第一个参数
type para372 struct {
	a int
	b []int
}

// ans 是答案
// one 代表第一个答案
type ans372 struct {
	one int
}

func Test_Problem372(t *testing.T) {

	qs := []question372{

		question372{
			para372{2, []int{3}},
			ans372{8},
		},

		question372{
			para372{2, []int{1, 0}},
			ans372{1024},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 372------------------------\n")

	for _, q := range qs {
		_, p := q.ans372, q.para372
		fmt.Printf("【input】:%v       【output】:%v\n", p, superPow(p.a, p.b))
	}
	fmt.Printf("\n\n\n")
}
