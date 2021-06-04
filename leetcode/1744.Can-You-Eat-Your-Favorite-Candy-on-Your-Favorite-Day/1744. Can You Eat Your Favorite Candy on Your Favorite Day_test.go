package leetcode

import (
	"fmt"
	"testing"
)

type question1744 struct {
	para1744
	ans1744
}

// para 是参数
// one 代表第一个参数
type para1744 struct {
	candiesCount []int
	queries      [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1744 struct {
	one []bool
}

func Test_Problem1744(t *testing.T) {

	qs := []question1744{

		{
			para1744{[]int{7, 4, 5, 3, 8}, [][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 1000000000}}},
			ans1744{[]bool{true, false, true}},
		},

		{
			para1744{[]int{5, 2, 6, 4, 1}, [][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1}}},
			ans1744{[]bool{false, true, true, false, false}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1744------------------------\n")

	for _, q := range qs {
		_, p := q.ans1744, q.para1744
		fmt.Printf("【input】:%v       【output】:%v\n", p, canEat(p.candiesCount, p.queries))
	}
	fmt.Printf("\n\n\n")
}
