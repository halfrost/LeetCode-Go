package leetcode

import (
	"fmt"
	"testing"
)

type question394 struct {
	para394
	ans394
}

// para 是参数
// one 代表第一个参数
type para394 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans394 struct {
	one string
}

func Test_Problem394(t *testing.T) {

	qs := []question394{

		question394{
			para394{"10[a]"},
			ans394{"aaaaaaaaaa"},
		},

		question394{
			para394{"3[a]2[bc]"},
			ans394{"aaabcbc"},
		},

		question394{
			para394{"3[a2[c]]"},
			ans394{"accaccacc"},
		},

		question394{
			para394{"2[abc]3[cd]ef"},
			ans394{"abcabccdcdcdef"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 394------------------------\n")

	for _, q := range qs {
		_, p := q.ans394, q.para394
		fmt.Printf("【input】:%v       【output】:%v\n", p, decodeString(p.s))
	}
	fmt.Printf("\n\n\n")
}
