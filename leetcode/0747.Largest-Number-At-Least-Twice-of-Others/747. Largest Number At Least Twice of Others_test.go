package leetcode

import (
	"fmt"
	"testing"
)

type question747 struct {
	para747
	ans747
}

// para 是参数
// one 代表第一个参数
type para747 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans747 struct {
	one int
}

func Test_Problem747(t *testing.T) {

	qs := []question747{

		{
			para747{[]int{3, 6, 1, 0}},
			ans747{1},
		},

		{
			para747{[]int{1, 2, 3, 4}},
			ans747{-1},
		},

		{
			para747{[]int{1}},
			ans747{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 747------------------------\n")

	for _, q := range qs {
		_, p := q.ans747, q.para747
		fmt.Printf("【input】:%v       【output】:%v\n", p, dominantIndex(p.nums))
	}
	fmt.Printf("\n\n\n")
}
