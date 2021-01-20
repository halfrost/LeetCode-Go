package leetcode

import (
	"fmt"
	"testing"
)

type question1680 struct {
	para1680
	ans1680
}

// para 是参数
// one 代表第一个参数
type para1680 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1680 struct {
	one int
}

func Test_Problem1680(t *testing.T) {

	qs := []question1680{

		{
			para1680{1},
			ans1680{1},
		},

		{
			para1680{3},
			ans1680{27},
		},

		{
			para1680{12},
			ans1680{505379714},
		},

		{
			para1680{42},
			ans1680{727837408},
		},

		{
			para1680{24},
			ans1680{385951001},
		},

		{
			para1680{81},
			ans1680{819357292},
		},

		{
			para1680{66},
			ans1680{627730462},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1680------------------------\n")

	for _, q := range qs {
		_, p := q.ans1680, q.para1680
		fmt.Printf("【input】:%v       【output】:%v\n", p, concatenatedBinary(p.n))
	}
	fmt.Printf("\n\n\n")
}
