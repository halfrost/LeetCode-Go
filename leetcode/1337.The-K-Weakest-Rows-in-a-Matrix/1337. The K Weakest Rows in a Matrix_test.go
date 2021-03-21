package leetcode

import (
	"fmt"
	"testing"
)

type question1337 struct {
	para1337
	ans1337
}

// para 是参数
// one 代表第一个参数
type para1337 struct {
	mat [][]int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans1337 struct {
	one []int
}

func Test_Problem1337(t *testing.T) {

	qs := []question1337{

		{
			para1337{[][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3},
			ans1337{[]int{2, 0, 3}},
		},

		{
			para1337{[][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 2},
			ans1337{[]int{0, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1337------------------------\n")

	for _, q := range qs {
		_, p := q.ans1337, q.para1337
		fmt.Printf("【input】:%v       【output】:%v\n", p, kWeakestRows(p.mat, p.k))
	}
	fmt.Printf("\n\n\n")
}
