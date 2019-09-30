package leetcode

import (
	"fmt"
	"testing"
)

type question154 struct {
	para154
	ans154
}

// para 是参数
// one 代表第一个参数
type para154 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans154 struct {
	one int
}

func Test_Problem154(t *testing.T) {

	qs := []question154{

		question154{
			para154{[]int{10, 2, 10, 10, 10, 10}},
			ans154{2},
		},

		question154{
			para154{[]int{1, 1}},
			ans154{1},
		},

		question154{
			para154{[]int{2, 2, 2, 0, 1}},
			ans154{0},
		},

		question154{
			para154{[]int{5, 1, 2, 3, 4}},
			ans154{1},
		},

		question154{
			para154{[]int{1}},
			ans154{1},
		},

		question154{
			para154{[]int{1, 2}},
			ans154{1},
		},

		question154{
			para154{[]int{2, 1}},
			ans154{1},
		},

		question154{
			para154{[]int{2, 3, 1}},
			ans154{1},
		},

		question154{
			para154{[]int{1, 2, 3}},
			ans154{1},
		},

		question154{
			para154{[]int{3, 4, 5, 1, 2}},
			ans154{1},
		},

		question154{
			para154{[]int{4, 5, 6, 7, 0, 1, 2}},
			ans154{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 154------------------------\n")

	for _, q := range qs {
		_, p := q.ans154, q.para154
		fmt.Printf("【input】:%v    【output】:%v\n", p, findMin154(p.nums))
	}
	fmt.Printf("\n\n\n")
}
