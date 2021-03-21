package leetcode

import (
	"fmt"
	"testing"
)

type question823 struct {
	para823
	ans823
}

// para 是参数
// one 代表第一个参数
type para823 struct {
	arr []int
}

// ans 是答案
// one 代表第一个答案
type ans823 struct {
	one int
}

func Test_Problem823(t *testing.T) {

	qs := []question823{

		{
			para823{[]int{2, 4}},
			ans823{3},
		},

		{
			para823{[]int{2, 4, 5, 10}},
			ans823{7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 823------------------------\n")

	for _, q := range qs {
		_, p := q.ans823, q.para823
		fmt.Printf("【input】:%v       【output】:%v\n", p, numFactoredBinaryTrees(p.arr))
	}
	fmt.Printf("\n\n\n")
}
