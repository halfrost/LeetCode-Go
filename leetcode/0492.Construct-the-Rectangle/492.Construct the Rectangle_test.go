package leetcode

import (
	"fmt"
	"testing"
)

type question492 struct {
	para492
	ans492
}

// area 是参数
type para492 struct {
	area int
}

// ans 是答案
type ans492 struct {
	ans []int
}

func Test_Problem492(t *testing.T) {

	qs := []question492{

		{
			para492{4},
			ans492{[]int{2, 2}},
		},

		{
			para492{37},
			ans492{[]int{37, 1}},
		},

		{
			para492{122122},
			ans492{[]int{427, 286}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 492------------------------\n")

	for _, q := range qs {
		_, p := q.ans492, q.para492
		fmt.Printf("【input】:%v       【output】:%v\n", p, constructRectangle(p.area))
	}
	fmt.Printf("\n\n\n")
}
