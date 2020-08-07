package leetcode

import (
	"fmt"
	"testing"
)

type question318 struct {
	para318
	ans318
}

// para 是参数
// one 代表第一个参数
type para318 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans318 struct {
	one int
}

func Test_Problem318(t *testing.T) {

	qs := []question318{

		question318{
			para318{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			ans318{16},
		},

		question318{
			para318{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			ans318{4},
		},

		question318{
			para318{[]string{"a", "aa", "aaa", "aaaa"}},
			ans318{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 318------------------------\n")

	for _, q := range qs {
		_, p := q.ans318, q.para318
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxProduct318(p.one))
	}
	fmt.Printf("\n\n\n")
}
