package leetcode

import (
	"fmt"
	"testing"
)

type question352 struct {
	para352
	ans352
}

// para 是参数
type para352 struct {
	para string
	num  int
}

// ans 是答案
type ans352 struct {
	ans [][]int
}

func Test_Problem352(t *testing.T) {

	qs := []question352{

		{
			para352{"addNum", 1},
			ans352{[][]int{{1, 1}}},
		},

		{
			para352{"addNum", 3},
			ans352{[][]int{{1, 1}, {3, 3}}},
		},

		{
			para352{"addNum", 7},
			ans352{[][]int{{1, 1}, {3, 3}, {7, 7}}},
		},

		{
			para352{"addNum", 2},
			ans352{[][]int{{1, 3}, {7, 7}}},
		},

		{
			para352{"addNum", 6},
			ans352{[][]int{{1, 3}, {6, 7}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 352------------------------\n")

	obj := Constructor()
	for _, q := range qs {
		_, p := q.ans352, q.para352
		obj.AddNum(p.num)
		fmt.Printf("【input】:%v       【output】:%v\n", p, obj.GetIntervals())
	}
	fmt.Printf("\n\n\n")
}
