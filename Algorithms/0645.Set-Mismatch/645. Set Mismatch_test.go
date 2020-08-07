package leetcode

import (
	"fmt"
	"testing"
)

type question645 struct {
	para645
	ans645
}

// para 是参数
// one 代表第一个参数
type para645 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans645 struct {
	one []int
}

func Test_Problem645(t *testing.T) {

	qs := []question645{

		question645{
			para645{[]int{1, 2, 2, 4}},
			ans645{[]int{2, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 645------------------------\n")

	for _, q := range qs {
		_, p := q.ans645, q.para645
		fmt.Printf("【input】:%v       【output】:%v\n", p, findErrorNums(p.one))
	}
	fmt.Printf("\n\n\n")
}
