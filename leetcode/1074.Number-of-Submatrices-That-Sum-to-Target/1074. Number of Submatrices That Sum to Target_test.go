package leetcode

import (
	"fmt"
	"testing"
)

type question1074 struct {
	para1074
	ans1074
}

// para 是参数
// one 代表第一个参数
type para1074 struct {
	one [][]int
	t   int
}

// ans 是答案
// one 代表第一个答案
type ans1074 struct {
	one int
}

func Test_Problem1074(t *testing.T) {

	qs := []question1074{

		question1074{
			para1074{[][]int{[]int{0, 1, 0}, []int{1, 1, 1}, []int{0, 1, 0}}, 0},
			ans1074{4},
		},

		question1074{
			para1074{[][]int{[]int{1, -1}, []int{-1, 1}}, 0},
			ans1074{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1074------------------------\n")

	for _, q := range qs {
		_, p := q.ans1074, q.para1074
		fmt.Printf("【input】:%v       【output】:%v\n", p, numSubmatrixSumTarget1(p.one, p.t))
	}
	fmt.Printf("\n\n\n")
}
