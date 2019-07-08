package leetcode

import (
	"fmt"
	"testing"
)

type question213 struct {
	para213
	ans213
}

// para 是参数
// one 代表第一个参数
type para213 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans213 struct {
	one int
}

func Test_Problem213(t *testing.T) {

	qs := []question213{

		question213{
			para213{[]int{0, 0}},
			ans213{0},
		},

		question213{
			para213{[]int{2, 3, 2}},
			ans213{3},
		},
		question213{
			para213{[]int{1, 2, 3, 1}},
			ans213{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 213------------------------\n")

	for _, q := range qs {
		_, p := q.ans213, q.para213
		fmt.Printf("【input】:%v       【output】:%v\n", p, rob213(p.one))
	}
	fmt.Printf("\n\n\n")
}
