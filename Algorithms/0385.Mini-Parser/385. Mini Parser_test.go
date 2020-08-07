package leetcode

import (
	"fmt"
	"testing"
)

type question385 struct {
	para385
	ans385
}

// para 是参数
// one 代表第一个参数
type para385 struct {
	n string
}

// ans 是答案
// one 代表第一个答案
type ans385 struct {
	one []int
}

func Test_Problem385(t *testing.T) {

	qs := []question385{

		question385{
			para385{"[[]]"},
			ans385{[]int{}},
		},

		question385{
			para385{"[]"},
			ans385{[]int{}},
		},

		question385{
			para385{"[-1]"},
			ans385{[]int{-1}},
		},

		question385{
			para385{"[123,[456,[789]]]"},
			ans385{[]int{123, 456, 789}},
		},

		question385{
			para385{"324"},
			ans385{[]int{324}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 385------------------------\n")

	for _, q := range qs {
		_, p := q.ans385, q.para385
		fmt.Printf("【input】:%v       【output】: \n", p)
		fmt.Printf("NestedInteger = ")
		deserialize(p.n).Print()
	}
	fmt.Printf("\n\n\n")
}
