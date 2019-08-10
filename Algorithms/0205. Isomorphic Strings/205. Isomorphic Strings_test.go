package leetcode

import (
	"fmt"
	"testing"
)

type question205 struct {
	para205
	ans205
}

// para 是参数
// one 代表第一个参数
type para205 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans205 struct {
	one bool
}

func Test_Problem205(t *testing.T) {

	qs := []question205{

		question205{
			para205{"egg", "add"},
			ans205{true},
		},

		question205{
			para205{"foo", "bar"},
			ans205{false},
		},

		question205{
			para205{"paper", "title"},
			ans205{true},
		},

		question205{
			para205{"", ""},
			ans205{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 205------------------------\n")

	for _, q := range qs {
		_, p := q.ans205, q.para205
		fmt.Printf("【input】:%v       【output】:%v\n", p, isIsomorphic(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
