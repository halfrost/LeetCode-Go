package leetcode

import (
	"fmt"
	"testing"
)

type question438 struct {
	para438
	ans438
}

// para 是参数
// one 代表第一个参数
type para438 struct {
	s string
	p string
}

// ans 是答案
// one 代表第一个答案
type ans438 struct {
	one []int
}

func Test_Problem438(t *testing.T) {

	qs := []question438{

		question438{
			para438{"abab", "ab"},
			ans438{[]int{0, 1, 2}},
		},

		question438{
			para438{"cbaebabacd", "abc"},
			ans438{[]int{0, 6}},
		},

		question438{
			para438{"", "abc"},
			ans438{[]int{}},
		},

		question438{
			para438{"abacbabc", "abc"},
			ans438{[]int{1, 2, 3, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 438------------------------\n")

	for _, q := range qs {
		_, p := q.ans438, q.para438
		fmt.Printf("【input】:%v       【output】:%v\n", p, findAnagrams(p.s, p.p))
	}
	fmt.Printf("\n\n\n")
}
