package leetcode

import (
	"fmt"
	"testing"
)

type question414 struct {
	para414
	ans414
}

// para 是参数
// one 代表第一个参数
type para414 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans414 struct {
	one int
}

func Test_Problem414(t *testing.T) {

	qs := []question414{

		question414{
			para414{[]int{1, 1, 2}},
			ans414{2},
		},

		question414{
			para414{[]int{3, 2, 1}},
			ans414{1},
		},

		question414{
			para414{[]int{1, 2}},
			ans414{2},
		},

		question414{
			para414{[]int{2, 2, 3, 1}},
			ans414{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 414------------------------\n")

	for _, q := range qs {
		_, p := q.ans414, q.para414
		fmt.Printf("【input】:%v       【output】:%v\n", p, thirdMax(p.one))
	}
	fmt.Printf("\n\n\n")
}
