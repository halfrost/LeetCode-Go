package leetcode

import (
	"fmt"
	"testing"
)

type question970 struct {
	para970
	ans970
}

// para 是参数
// one 代表第一个参数
type para970 struct {
	one int
	two int
	b   int
}

// ans 是答案
// one 代表第一个答案
type ans970 struct {
	one []int
}

func Test_Problem970(t *testing.T) {

	qs := []question970{

		question970{
			para970{2, 3, 10},
			ans970{[]int{2, 3, 4, 5, 7, 9, 10}},
		},

		question970{
			para970{3, 5, 15},
			ans970{[]int{2, 4, 6, 8, 10, 14}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 970------------------------\n")

	for _, q := range qs {
		_, p := q.ans970, q.para970
		fmt.Printf("【input】:%v       【output】:%v\n", p, powerfulIntegers(p.one, p.two, p.b))
	}
	fmt.Printf("\n\n\n")
}
