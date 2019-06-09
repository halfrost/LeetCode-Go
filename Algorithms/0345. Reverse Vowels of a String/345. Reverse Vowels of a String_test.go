package leetcode

import (
	"fmt"
	"testing"
)

type question345 struct {
	para345
	ans345
}

// para 是参数
// one 代表第一个参数
type para345 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans345 struct {
	one string
}

func Test_Problem345(t *testing.T) {

	qs := []question345{

		question345{
			para345{"hello"},
			ans345{"holle"},
		},

		question345{
			para345{"leetcode"},
			ans345{"leotcede"},
		},

		question345{
			para345{"aA"},
			ans345{"Aa"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 345------------------------\n")

	for _, q := range qs {
		_, p := q.ans345, q.para345
		fmt.Printf("【input】:%v       【output】:%v\n", p, reverseVowels(p.one))
	}
	fmt.Printf("\n\n\n")
}
