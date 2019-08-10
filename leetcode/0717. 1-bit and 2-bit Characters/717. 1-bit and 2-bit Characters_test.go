package leetcode

import (
	"fmt"
	"testing"
)

type question717 struct {
	para717
	ans717
}

// para 是参数
// one 代表第一个参数
type para717 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans717 struct {
	one bool
}

func Test_Problem717(t *testing.T) {

	qs := []question717{

		question717{
			para717{[]int{1, 0, 0}},
			ans717{true},
		},

		question717{
			para717{[]int{1, 1, 1, 0}},
			ans717{false},
		},

		question717{
			para717{[]int{0, 1, 1, 1, 0, 0}},
			ans717{true},
		},

		question717{
			para717{[]int{1, 1, 1, 1, 0}},
			ans717{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 717------------------------\n")

	for _, q := range qs {
		_, p := q.ans717, q.para717
		fmt.Printf("【input】:%v       【output】:%v\n", p, isOneBitCharacter(p.one))
	}
	fmt.Printf("\n\n\n")
}
