package leetcode

import (
	"fmt"
	"testing"
)

type question524 struct {
	para524
	ans524
}

// para 是参数
// one 代表第一个参数
type para524 struct {
	s   string
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans524 struct {
	one string
}

func Test_Problem524(t *testing.T) {

	qs := []question524{

		question524{
			para524{"abpcplea", []string{"ale", "apple", "monkey", "plea"}},
			ans524{"apple"},
		},

		question524{
			para524{"abpcplea", []string{"a", "b", "c"}},
			ans524{"a"},
		},

		question524{
			para524{"abpcplea", []string{"aaaaa", "b", "c"}},
			ans524{"b"},
		},

		question524{
			para524{"bab", []string{"ba", "ab", "a", "b"}},
			ans524{"ab"},
		},

		question524{
			para524{"aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}},
			ans524{"ewaf"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 524------------------------\n")

	for _, q := range qs {
		_, p := q.ans524, q.para524
		fmt.Printf("【input】:%v       【output】:%v\n", p, findLongestWord(p.s, p.one))
	}
	fmt.Printf("\n\n\n")
}
