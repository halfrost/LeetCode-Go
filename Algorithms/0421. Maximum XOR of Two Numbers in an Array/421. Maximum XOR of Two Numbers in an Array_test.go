package leetcode

import (
	"fmt"
	"testing"
)

type question421 struct {
	para421
	ans421
}

// para 是参数
// one 代表第一个参数
type para421 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans421 struct {
	one int
}

func Test_Problem421(t *testing.T) {

	qs := []question421{

		question421{
			para421{[]int{3, 10, 5, 25, 2, 8}},
			ans421{28},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 421------------------------\n")

	for _, q := range qs {
		_, p := q.ans421, q.para421
		fmt.Printf("【input】:%v       【output】:%v\n", p, findMaximumXOR(p.one))
	}
	fmt.Printf("\n\n\n")
}
