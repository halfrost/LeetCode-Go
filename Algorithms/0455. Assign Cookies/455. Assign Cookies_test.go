package leetcode

import (
	"fmt"
	"testing"
)

type question455 struct {
	para455
	ans455
}

// para 是参数
// one 代表第一个参数
type para455 struct {
	g []int
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans455 struct {
	one int
}

func Test_Problem455(t *testing.T) {

	qs := []question455{

		question455{
			para455{[]int{1, 2, 3}, []int{1, 1}},
			ans455{1},
		},

		question455{
			para455{[]int{1, 2}, []int{1, 2, 3}},
			ans455{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 455------------------------\n")

	for _, q := range qs {
		_, p := q.ans455, q.para455
		fmt.Printf("【input】:%v       【output】:%v\n", p, findContentChildren(p.g, p.s))
	}
	fmt.Printf("\n\n\n")
}
