package leetcode

import (
	"fmt"
	"testing"
)

type question547 struct {
	para547
	ans547
}

// para 是参数
// one 代表第一个参数
type para547 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans547 struct {
	one int
}

func Test_Problem547(t *testing.T) {

	qs := []question547{

		question547{
			para547{[][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}},
			ans547{3},
		},

		question547{
			para547{[][]int{[]int{1, 1, 0}, []int{1, 1, 0}, []int{0, 0, 1}}},
			ans547{2},
		},

		question547{
			para547{[][]int{[]int{1, 1, 0}, []int{1, 1, 1}, []int{0, 1, 1}}},
			ans547{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 547------------------------\n")

	for _, q := range qs {
		_, p := q.ans547, q.para547
		fmt.Printf("【input】:%v       【output】:%v\n", p, findCircleNum(p.one))
	}
	fmt.Printf("\n\n\n")
}
