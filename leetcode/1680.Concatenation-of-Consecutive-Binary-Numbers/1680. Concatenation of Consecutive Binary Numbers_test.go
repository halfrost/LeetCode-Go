package leetcode

import (
	"fmt"
	"testing"
)

type question5620 struct {
	para5620
	ans5620
}

// para 是参数
// one 代表第一个参数
type para5620 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans5620 struct {
	one int
}

func Test_Problem5620(t *testing.T) {

	qs := []question5620{

		{
			para5620{1},
			ans5620{1},
		},

		{
			para5620{3},
			ans5620{27},
		},

		{
			para5620{12},
			ans5620{505379714},
		},

		{
			para5620{42},
			ans5620{727837408},
		},

		{
			para5620{24},
			ans5620{385951001},
		},

		{
			para5620{81},
			ans5620{819357292},
		},

		{
			para5620{66},
			ans5620{627730462},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 5620------------------------\n")

	for _, q := range qs {
		_, p := q.ans5620, q.para5620
		fmt.Printf("【input】:%v       【output】:%v\n", p, concatenatedBinary(p.n))
	}
	fmt.Printf("\n\n\n")
}
