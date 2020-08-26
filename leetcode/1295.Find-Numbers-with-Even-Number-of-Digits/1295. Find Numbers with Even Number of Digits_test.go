package leetcode

import (
	"fmt"
	"testing"
)

type question1295 struct {
	para1295
	ans1295
}

// para 是参数
// one 代表第一个参数
type para1295 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1295 struct {
	one int
}

func Test_Problem1295(t *testing.T) {

	qs := []question1295{

		{
			para1295{[]int{12, 345, 2, 6, 7896}},
			ans1295{2},
		},

		{
			para1295{[]int{555, 901, 482, 1771}},
			ans1295{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1295------------------------\n")

	for _, q := range qs {
		_, p := q.ans1295, q.para1295
		fmt.Printf("【input】:%v       【output】:%v\n", p, findNumbers(p.one))
	}
	fmt.Printf("\n\n\n")
}
