package leetcode

import (
	"fmt"
	"testing"
)

type question290 struct {
	para290
	ans290
}

// para 是参数
// one 代表第一个参数
type para290 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans290 struct {
	one bool
}

func Test_Problem290(t *testing.T) {

	qs := []question290{

		question290{
			para290{"abba", "dog cat cat dog"},
			ans290{true},
		},

		question290{
			para290{"abba", "dog cat cat fish"},
			ans290{false},
		},

		question290{
			para290{"aaaa", "dog cat cat dog"},
			ans290{false},
		},

		question290{
			para290{"abba", "dog dog dog dog"},
			ans290{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 290------------------------\n")

	for _, q := range qs {
		_, p := q.ans290, q.para290
		fmt.Printf("【input】:%v       【output】:%v\n", p, wordPattern(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
