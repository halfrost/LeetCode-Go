package leetcode

import (
	"fmt"
	"testing"
)

type question795 struct {
	para795
	ans795
}

// para 是参数
// one 代表第一个参数
type para795 struct {
	nums  []int
	left  int
	right int
}

// ans 是答案
// one 代表第一个答案
type ans795 struct {
	one int
}

func Test_Problem795(t *testing.T) {

	qs := []question795{

		{
			para795{[]int{2, 1, 4, 3}, 2, 3},
			ans795{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 795------------------------\n")

	for _, q := range qs {
		_, p := q.ans795, q.para795
		fmt.Printf("【input】:%v       【output】:%v\n", p, numSubarrayBoundedMax(p.nums, p.left, p.right))
	}
	fmt.Printf("\n\n\n")
}
