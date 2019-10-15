package leetcode

import (
	"fmt"
	"testing"
)

type question633 struct {
	para633
	ans633
}

// para 是参数
// one 代表第一个参数
type para633 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans633 struct {
	one bool
}

func Test_Problem633(t *testing.T) {

	qs := []question633{

		question633{
			para633{1},
			ans633{true},
		},

		question633{
			para633{2},
			ans633{true},
		},

		question633{
			para633{3},
			ans633{false},
		},

		question633{
			para633{4},
			ans633{true},
		},

		question633{
			para633{5},
			ans633{true},
		},

		question633{
			para633{6},
			ans633{false},
		},

		question633{
			para633{104976},
			ans633{true},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 633------------------------\n")

	for _, q := range qs {
		_, p := q.ans633, q.para633
		fmt.Printf("【input】:%v       【output】:%v\n", p, judgeSquareSum(p.one))
	}
	fmt.Printf("\n\n\n")
}
