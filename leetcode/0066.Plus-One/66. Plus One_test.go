package leetcode

import (
	"fmt"
	"testing"
)

type question66 struct {
	para66
	ans66
}

// para 是参数
// one 代表第一个参数
type para66 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans66 struct {
	one []int
}

func Test_Problem66(t *testing.T) {

	qs := []question66{

		question66{
			para66{[]int{1, 2, 3}},
			ans66{[]int{1, 2, 4}},
		},

		question66{
			para66{[]int{4, 3, 2, 1}},
			ans66{[]int{4, 3, 2, 2}},
		},

		question66{
			para66{[]int{9, 9}},
			ans66{[]int{1, 0, 0}},
		},

		question66{
			para66{[]int{0}},
			ans66{[]int{0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 66------------------------\n")

	for _, q := range qs {
		_, p := q.ans66, q.para66
		fmt.Printf("【input】:%v       【output】:%v\n", p, plusOne(p.one))
	}
	fmt.Printf("\n\n\n")
}
