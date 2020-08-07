package leetcode

import (
	"fmt"
	"testing"
)

type question93 struct {
	para93
	ans93
}

// para 是参数
// one 代表第一个参数
type para93 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans93 struct {
	one []string
}

func Test_Problem93(t *testing.T) {

	qs := []question93{

		question93{
			para93{"25525511135"},
			ans93{[]string{"255.255.11.135", "255.255.111.35"}},
		},

		question93{
			para93{"0000"},
			ans93{[]string{"0.0.0.0"}},
		},

		question93{
			para93{"010010"},
			ans93{[]string{"0.10.0.10", "0.100.1.0"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 93------------------------\n")

	for _, q := range qs {
		_, p := q.ans93, q.para93
		fmt.Printf("【input】:%v       【output】:%v\n", p, restoreIPAddresses(p.s))
	}
	fmt.Printf("\n\n\n")
}
