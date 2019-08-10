package leetcode

import (
	"fmt"
	"testing"
)

type question207 struct {
	para207
	ans207
}

// para 是参数
// one 代表第一个参数
type para207 struct {
	one int
	pre [][]int
}

// ans 是答案
// one 代表第一个答案
type ans207 struct {
	one bool
}

func Test_Problem207(t *testing.T) {

	qs := []question207{

		question207{
			para207{2, [][]int{[]int{1, 0}}},
			ans207{true},
		},

		question207{
			para207{2, [][]int{[]int{1, 0}, []int{0, 1}}},
			ans207{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 207------------------------\n")

	for _, q := range qs {
		_, p := q.ans207, q.para207
		fmt.Printf("【input】:%v       【output】:%v\n", p, canFinish(p.one, p.pre))
	}
	fmt.Printf("\n\n\n")
}
