package leetcode

import (
	"fmt"
	"testing"
)

type question1105 struct {
	para1105
	ans1105
}

// para 是参数
// one 代表第一个参数
type para1105 struct {
	one [][]int
	w   int
}

// ans 是答案
// one 代表第一个答案
type ans1105 struct {
	one int
}

func Test_Problem1105(t *testing.T) {

	qs := []question1105{

		question1105{
			para1105{[][]int{[]int{1, 1}, []int{2, 3}, []int{2, 3}, []int{1, 1}, []int{1, 1}, []int{1, 1}, []int{1, 2}}, 4},
			ans1105{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1105------------------------\n")

	for _, q := range qs {
		_, p := q.ans1105, q.para1105
		fmt.Printf("【input】:%v       【output】:%v\n", p, minHeightShelves(p.one, p.w))
	}
	fmt.Printf("\n\n\n")
}
