package leetcode

import (
	"fmt"
	"testing"
)

type question685 struct {
	para685
	ans685
}

// para 是参数
// one 代表第一个参数
type para685 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans685 struct {
	one []int
}

func Test_Problem685(t *testing.T) {

	qs := []question685{

		question685{
			para685{[][]int{[]int{3, 5}, []int{1, 3}, []int{2, 1}, []int{5, 4}, []int{2, 3}}},
			ans685{[]int{2, 3}},
		},

		question685{
			para685{[][]int{[]int{4, 2}, []int{1, 5}, []int{5, 2}, []int{5, 3}, []int{2, 4}}},
			ans685{[]int{4, 2}},
		},

		question685{
			para685{[][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}},
			ans685{[]int{2, 3}},
		},

		question685{
			para685{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 4}, []int{1, 5}}},
			ans685{[]int{1, 4}},
		},

		question685{
			para685{[][]int{[]int{2, 1}, []int{3, 1}, []int{4, 2}, []int{1, 4}}},
			ans685{[]int{2, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 685------------------------\n")

	for _, q := range qs {
		_, p := q.ans685, q.para685
		fmt.Printf("【input】:%v       【output】:%v\n", p, findRedundantDirectedConnection(p.one))
	}
	fmt.Printf("\n\n\n")
}
