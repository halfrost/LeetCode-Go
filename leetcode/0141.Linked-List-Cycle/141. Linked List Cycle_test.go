package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question141 struct {
	para141
	ans141
}

// para 是参数
// one 代表第一个参数
type para141 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans141 struct {
	one bool
}

func Test_Problem141(t *testing.T) {

	qs := []question141{

		{
			para141{[]int{3, 2, 0, -4}},
			ans141{false},
		},

		{
			para141{[]int{1, 2}},
			ans141{false},
		},

		{
			para141{[]int{1}},
			ans141{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 141------------------------\n")

	for _, q := range qs {
		_, p := q.ans141, q.para141
		fmt.Printf("【input】:%v       【output】:%v\n", p, hasCycle(structures.Ints2List(p.one)))
	}
	fmt.Printf("\n\n\n")
}
