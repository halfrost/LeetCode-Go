package leetcode

import (
	"fmt"
	"testing"
)

type question151 struct {
	para151
	ans151
}

// para 是参数
// one 代表第一个参数
type para151 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans151 struct {
	one string
}

func Test_Problem151(t *testing.T) {

	qs := []question151{

		question151{
			para151{"the sky is blue"},
			ans151{"blue is sky the"},
		},

		question151{
			para151{"  hello world!  "},
			ans151{"world! hello"},
		},
		question151{
			para151{"a good   example"},
			ans151{"example good a"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 151------------------------\n")

	for _, q := range qs {
		_, p := q.ans151, q.para151
		fmt.Printf("【input】:%v       【output】:%v\n", p, reverseWords151(p.one))
	}
	fmt.Printf("\n\n\n")
}
