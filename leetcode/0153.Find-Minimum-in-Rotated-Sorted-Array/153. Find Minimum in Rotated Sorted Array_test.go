package leetcode

import (
	"fmt"
	"testing"
)

type question153 struct {
	para153
	ans153
}

// para 是参数
// one 代表第一个参数
type para153 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans153 struct {
	one int
}

func Test_Problem153(t *testing.T) {

	qs := []question153{

		question153{
			para153{[]int{5, 1, 2, 3, 4}},
			ans153{1},
		},

		question153{
			para153{[]int{1}},
			ans153{1},
		},

		question153{
			para153{[]int{1, 2}},
			ans153{1},
		},

		question153{
			para153{[]int{2, 1}},
			ans153{1},
		},

		question153{
			para153{[]int{2, 3, 1}},
			ans153{1},
		},

		question153{
			para153{[]int{1, 2, 3}},
			ans153{1},
		},

		question153{
			para153{[]int{3, 4, 5, 1, 2}},
			ans153{1},
		},

		question153{
			para153{[]int{4, 5, 6, 7, 0, 1, 2}},
			ans153{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 153------------------------\n")

	for _, q := range qs {
		_, p := q.ans153, q.para153
		fmt.Printf("【input】:%v    【output】:%v\n", p, findMin(p.nums))
	}
	fmt.Printf("\n\n\n")
}
