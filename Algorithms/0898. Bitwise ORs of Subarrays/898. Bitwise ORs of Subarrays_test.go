package leetcode

import (
	"fmt"
	"testing"
)

type question898 struct {
	para898
	ans898
}

// para 是参数
// one 代表第一个参数
type para898 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans898 struct {
	one int
}

func Test_Problem898(t *testing.T) {

	qs := []question898{

		question898{
			para898{[]int{0}},
			ans898{1},
		},

		question898{
			para898{[]int{1, 1, 2}},
			ans898{3},
		},

		question898{
			para898{[]int{1, 2, 4}},
			ans898{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 898------------------------\n")

	for _, q := range qs {
		_, p := q.ans898, q.para898
		fmt.Printf("【input】:%v       【output】:%v\n", p, subarrayBitwiseORs(p.one))
	}
	fmt.Printf("\n\n\n")
}
