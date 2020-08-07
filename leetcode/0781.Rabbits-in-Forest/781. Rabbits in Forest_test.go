package leetcode

import (
	"fmt"
	"testing"
)

type question781 struct {
	para781
	ans781
}

// para 是参数
// one 代表第一个参数
type para781 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans781 struct {
	one int
}

func Test_Problem781(t *testing.T) {

	qs := []question781{
		question781{
			para781{[]int{1, 1, 2}},
			ans781{5},
		},

		question781{
			para781{[]int{10, 10, 10}},
			ans781{11},
		},

		question781{
			para781{[]int{}},
			ans781{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 781------------------------\n")

	for _, q := range qs {
		_, p := q.ans781, q.para781
		fmt.Printf("【input】:%v       【output】:%v\n", p, numRabbits(p.one))
	}
	fmt.Printf("\n\n\n")
}
