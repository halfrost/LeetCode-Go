package leetcode

import (
	"fmt"
	"testing"
)

type question119 struct {
	para119
	ans119
}

// para 是参数
// one 代表第一个参数
type para119 struct {
	rowIndex int
}

// ans 是答案
// one 代表第一个答案
type ans119 struct {
	one []int
}

func Test_Problem119(t *testing.T) {

	qs := []question119{

		{
			para119{3},
			ans119{[]int{1, 3, 3, 1}},
		},

		{
			para119{0},
			ans119{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 119------------------------\n")

	for _, q := range qs {
		_, p := q.ans119, q.para119
		fmt.Printf("【input】:%v    【output】:%v\n", p, getRow(p.rowIndex))
	}
	fmt.Printf("\n\n\n")
}
