package leetcode

import (
	"fmt"
	"testing"
)

type question541 struct {
	para541
	ans541
}

// para 是参数
// one 代表第一个参数
type para541 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans541 struct {
	one string
}

func Test_Problem541(t *testing.T) {

	qs := []question541{

		question541{
			para541{"abcdefg", 2},
			ans541{"bacdfeg"},
		},

		question541{
			para541{"abcdefg", 5},
			ans541{"edcbafg"},
		},

		question541{
			para541{"abcd", 4},
			ans541{"dcba"},
		},

		question541{
			para541{"", 100},
			ans541{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 541------------------------\n")

	for _, q := range qs {
		_, p := q.ans541, q.para541
		fmt.Printf("【input】:%v       【output】:%v\n", p, reverseStr(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
