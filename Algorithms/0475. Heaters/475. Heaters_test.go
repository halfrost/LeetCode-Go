package leetcode

import (
	"fmt"
	"testing"
)

type question475 struct {
	para475
	ans475
}

// para 是参数
// one 代表第一个参数
type para475 struct {
	houses  []int
	heaters []int
}

// ans 是答案
// one 代表第一个答案
type ans475 struct {
	one int
}

func Test_Problem475(t *testing.T) {

	qs := []question475{

		question475{
			para475{[]int{1, 2, 3}, []int{2}},
			ans475{1},
		},

		question475{
			para475{[]int{1, 2, 3, 4}, []int{1, 4}},
			ans475{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 475------------------------\n")

	for _, q := range qs {
		_, p := q.ans475, q.para475
		fmt.Printf("【input】:%v       【output】:%v\n", p, findRadius(p.houses, p.heaters))
	}
	fmt.Printf("\n\n\n")
}
