package leetcode

import (
	"fmt"
	"testing"
)

type question995 struct {
	para995
	ans995
}

// para 是参数
// one 代表第一个参数
type para995 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans995 struct {
	one int
}

func Test_Problem995(t *testing.T) {

	qs := []question995{

		question995{
			para995{[]int{0, 1, 0}, 1},
			ans995{2},
		},

		question995{
			para995{[]int{1, 1, 0}, 2},
			ans995{-1},
		},

		question995{
			para995{[]int{0, 0, 0, 1, 0, 1, 1, 0}, 3},
			ans995{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 995------------------------\n")

	for _, q := range qs {
		_, p := q.ans995, q.para995
		fmt.Printf("【input】:%v       【output】:%v\n", p, minKBitFlips(p.one, p.k))
	}
	fmt.Printf("\n\n\n")
}
