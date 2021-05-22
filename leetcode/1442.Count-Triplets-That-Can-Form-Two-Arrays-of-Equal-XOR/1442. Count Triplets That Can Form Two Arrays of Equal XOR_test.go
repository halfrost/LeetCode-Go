package leetcode

import (
	"fmt"
	"testing"
)

type question1442 struct {
	para1442
	ans1442
}

// para 是参数
// one 代表第一个参数
type para1442 struct {
	arr []int
}

// ans 是答案
// one 代表第一个答案
type ans1442 struct {
	one int
}

func Test_Problem1442(t *testing.T) {

	qs := []question1442{

		{
			para1442{[]int{2, 3, 1, 6, 7}},
			ans1442{4},
		},

		{
			para1442{[]int{1, 1, 1, 1, 1}},
			ans1442{10},
		},

		{
			para1442{[]int{2, 3}},
			ans1442{0},
		},

		{
			para1442{[]int{1, 3, 5, 7, 9}},
			ans1442{3},
		},

		{
			para1442{[]int{7, 11, 12, 9, 5, 2, 7, 17, 22}},
			ans1442{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1442------------------------\n")

	for _, q := range qs {
		_, p := q.ans1442, q.para1442
		fmt.Printf("【input】:%v       【output】:%v\n", p, countTriplets(p.arr))
	}
	fmt.Printf("\n\n\n")
}
