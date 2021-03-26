package leetcode

import (
	"fmt"
	"testing"
)

type question870 struct {
	para870
	ans870
}

// para 是参数
// one 代表第一个参数
type para870 struct {
	A []int
	B []int
}

// ans 是答案
// one 代表第一个答案
type ans870 struct {
	one []int
}

func Test_Problem870(t *testing.T) {

	qs := []question870{

		{
			para870{[]int{2, 7, 11, 15}, []int{1, 10, 4, 11}},
			ans870{[]int{2, 11, 7, 15}},
		},

		{
			para870{[]int{12, 24, 8, 32}, []int{13, 25, 32, 11}},
			ans870{[]int{24, 32, 8, 12}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 870------------------------\n")

	for _, q := range qs {
		_, p := q.ans870, q.para870
		fmt.Printf("【input】:%v       【output】:%v\n", p, advantageCount1(p.A, p.B))
	}
	fmt.Printf("\n\n\n")
}
