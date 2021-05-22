package leetcode

import (
	"fmt"
	"testing"
)

type question1572 struct {
	para1572
	ans1572
}

// para 是参数
type para1572 struct {
	mat [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1572 struct {
	one int
}

func Test_Problem1572(t *testing.T) {

	qs := []question1572{

		{
			para1572{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			ans1572{25},
		},

		{
			para1572{[][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}},
			ans1572{8},
		},

		{
			para1572{[][]int{{5}}},
			ans1572{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1572------------------------\n")

	for _, q := range qs {
		_, p := q.ans1572, q.para1572
		fmt.Printf("【input】:%v      【output】:%v      \n", p, diagonalSum(p.mat))
	}
	fmt.Printf("\n\n\n")
}
