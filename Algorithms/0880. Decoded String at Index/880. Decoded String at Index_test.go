package leetcode

import (
	"fmt"
	"testing"
)

type question880 struct {
	para880
	ans880
}

// para 是参数
// one 代表第一个参数
type para880 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans880 struct {
	one string
}

func Test_Problem880(t *testing.T) {

	qs := []question880{

		question880{
			para880{"aw4eguc6cs", 41},
			ans880{"a"},
		},

		question880{
			para880{"vk6u5xhq9v", 554},
			ans880{"h"},
		},

		question880{
			para880{"leet2code3", 10},
			ans880{"o"},
		},

		question880{
			para880{"ha22", 5},
			ans880{"h"},
		},

		question880{
			para880{"a2345678999999999999999", 1},
			ans880{"a"},
		},

		question880{
			para880{"y959q969u3hb22odq595", 222280369},
			ans880{"q"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 880------------------------\n")

	for _, q := range qs {
		_, p := q.ans880, q.para880
		fmt.Printf("【input】:%v       【output】:%v\n", p, decodeAtIndex(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
