package leetcode

import (
	"fmt"
	"testing"
)

type question462 struct {
	para462
	ans462
}

// para 是参数
// one 代表第一个参数
type para462 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans462 struct {
	one int
}

func Test_Problem462(t *testing.T) {

	qs := []question462{

		{
			para462{[]int{}},
			ans462{0},
		},

		{
			para462{[]int{1, 2, 3}},
			ans462{2},
		},

		{
			para462{[]int{1, 10, 2, 9}},
			ans462{16},
		},

		{
			para462{[]int{1, 0, 0, 8, 6}},
			ans462{14},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 462------------------------\n")

	for _, q := range qs {
		_, p := q.ans462, q.para462
		fmt.Printf("【input】:%v       【output】:%v\n", p, minMoves2(p.nums))
	}
	fmt.Printf("\n\n\n")
}
