package leetcode

import (
	"fmt"
	"testing"
)

type question260 struct {
	para260
	ans260
}

// para 是参数
// one 代表第一个参数
type para260 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans260 struct {
	one []int
}

func Test_Problem260(t *testing.T) {

	qs := []question260{

		question260{
			para260{[]int{1, 2, 1, 3, 2, 5}},
			ans260{[]int{3, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 260------------------------\n")

	for _, q := range qs {
		_, p := q.ans260, q.para260
		fmt.Printf("【input】:%v       【output】:%v\n", p, singleNumberIII(p.s))
	}
	fmt.Printf("\n\n\n")
}
