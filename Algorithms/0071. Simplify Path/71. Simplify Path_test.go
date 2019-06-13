package leetcode

import (
	"fmt"
	"testing"
)

type question71 struct {
	para71
	ans71
}

// para 是参数
// one 代表第一个参数
type para71 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans71 struct {
	one string
}

func Test_Problem71(t *testing.T) {

	qs := []question71{

		question71{
			para71{"/.hidden"},
			ans71{"/.hidden"},
		},

		question71{
			para71{"/..hidden"},
			ans71{"/..hidden"},
		},

		question71{
			para71{"/abc/..."},
			ans71{"/abc/..."},
		},

		question71{
			para71{"/home/"},
			ans71{"/home"},
		},

		question71{
			para71{"/..."},
			ans71{"/..."},
		},

		question71{
			para71{"/../"},
			ans71{"/"},
		},

		question71{
			para71{"/home//foo/"},
			ans71{"/home/foo"},
		},

		question71{
			para71{"/a/./b/../../c/"},
			ans71{"/c"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 71------------------------\n")

	for _, q := range qs {
		_, p := q.ans71, q.para71
		fmt.Printf("【input】:%v       【output】:%v\n", p, simplifyPath(p.s))
	}
	fmt.Printf("\n\n\n")
}
