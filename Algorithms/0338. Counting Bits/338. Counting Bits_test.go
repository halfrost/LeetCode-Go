package leetcode

import (
	"fmt"
	"testing"
)

type question338 struct {
	para338
	ans338
}

// para 是参数
// one 代表第一个参数
type para338 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans338 struct {
	one []int
}

func Test_Problem338(t *testing.T) {

	qs := []question338{

		question338{
			para338{2},
			ans338{[]int{0, 1, 1}},
		},

		question338{
			para338{5},
			ans338{[]int{0, 1, 1, 2, 1, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 338------------------------\n")

	for _, q := range qs {
		_, p := q.ans338, q.para338
		fmt.Printf("【input】:%v       【output】:%v\n", p, countBits(p.one))
	}
	fmt.Printf("\n\n\n")
}
