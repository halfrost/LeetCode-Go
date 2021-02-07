package leetcode

import (
	"fmt"
	"testing"
)

type question821 struct {
	para821
	ans821
}

// para 是参数
// one 代表第一个参数
type para821 struct {
	s string
	c byte
}

// ans 是答案
// one 代表第一个答案
type ans821 struct {
	one []int
}

func Test_Problem821(t *testing.T) {

	qs := []question821{

		{
			para821{"loveleetcode", 'e'},
			ans821{[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		},

		{
			para821{"baaa", 'b'},
			ans821{[]int{0, 1, 2, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 821------------------------\n")

	for _, q := range qs {
		_, p := q.ans821, q.para821
		fmt.Printf("【input】:%v       【output】:%v\n", p, shortestToChar(p.s, p.c))
	}
	fmt.Printf("\n\n\n")
}
