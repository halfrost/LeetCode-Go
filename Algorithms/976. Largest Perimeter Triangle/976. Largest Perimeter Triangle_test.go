package leetcode

import (
	"fmt"
	"testing"
)

type question976 struct {
	para976
	ans976
}

// para 是参数
// one 代表第一个参数
type para976 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans976 struct {
	one int
}

func Test_Problem976(t *testing.T) {

	qs := []question976{

		question976{
			para976{[]int{1, 2}},
			ans976{0},
		},
		question976{
			para976{[]int{1, 2, 3}},
			ans976{0},
		},

		question976{
			para976{[]int{}},
			ans976{0},
		},

		question976{
			para976{[]int{2, 1, 2}},
			ans976{5},
		},

		question976{
			para976{[]int{1, 1, 2}},
			ans976{0},
		},

		question976{
			para976{[]int{3, 2, 3, 4}},
			ans976{10},
		},

		question976{
			para976{[]int{3, 6, 2, 3}},
			ans976{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 976------------------------\n")

	for _, q := range qs {
		_, p := q.ans976, q.para976
		fmt.Printf("【input】:%v       【output】:%v\n", p, largestPerimeter(p.one))
	}
	fmt.Printf("\n\n\n")
}
