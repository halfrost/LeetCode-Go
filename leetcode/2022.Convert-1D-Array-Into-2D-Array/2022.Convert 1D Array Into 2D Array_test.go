package leetcode

import (
	"fmt"
	"testing"
)

type question2022 struct {
	para2022
	ans2022
}

// para 是参数
type para2022 struct {
	original []int
	m        int
	n        int
}

// ans 是答案
type ans2022 struct {
	ans [][]int
}

func Test_Problem2022(t *testing.T) {

	qs := []question2022{

		{
			para2022{[]int{1, 2, 3, 4}, 2, 2},
			ans2022{[][]int{{1, 2}, {3, 4}}},
		},

		{
			para2022{[]int{1, 2, 3}, 1, 3},
			ans2022{[][]int{{1, 2, 3}}},
		},

		{
			para2022{[]int{1, 2}, 1, 1},
			ans2022{nil},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2022------------------------\n")

	for _, q := range qs {
		_, p := q.ans2022, q.para2022
		fmt.Printf("【input】:%v    【output】:%v\n", p, construct2DArray(p.original, p.m, p.n))
	}
	fmt.Printf("\n\n\n")
}
