package leetcode

import (
	"fmt"
	"testing"
)

type question834 struct {
	para834
	ans834
}

// para 是参数
// one 代表第一个参数
type para834 struct {
	N     int
	edges [][]int
}

// ans 是答案
// one 代表第一个答案
type ans834 struct {
	one []int
}

func Test_Problem834(t *testing.T) {

	qs := []question834{

		question834{
			para834{4, [][]int{[]int{1, 2}, []int{3, 2}, []int{3, 0}}},
			ans834{[]int{6, 6, 4, 4}},
		},

		question834{
			para834{6, [][]int{[]int{0, 1}, []int{0, 2}, []int{2, 3}, []int{2, 4}, []int{2, 5}}},
			ans834{[]int{8, 12, 6, 10, 10, 10}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 834------------------------\n")

	for _, q := range qs {
		_, p := q.ans834, q.para834
		fmt.Printf("【input】:%v       【output】:%v\n", p, sumOfDistancesInTree(p.N, p.edges))
	}
	fmt.Printf("\n\n\n")
}
