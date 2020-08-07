package leetcode

import (
	"fmt"
	"testing"
)

type question409 struct {
	para409
	ans409
}

// para 是参数
// one 代表第一个参数
type para409 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans409 struct {
	one int
}

func Test_Problem409(t *testing.T) {

	qs := []question409{

		question409{
			para409{"abccccdd"},
			ans409{7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 409------------------------\n")

	for _, q := range qs {
		_, p := q.ans409, q.para409
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestPalindrome(p.one))
	}
	fmt.Printf("\n\n\n")
}
