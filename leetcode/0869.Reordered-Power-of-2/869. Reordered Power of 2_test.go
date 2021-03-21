package leetcode

import (
	"fmt"
	"testing"
)

type question869 struct {
	para869
	ans869
}

// para 是参数
// one 代表第一个参数
type para869 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans869 struct {
	one bool
}

func Test_Problem869(t *testing.T) {

	qs := []question869{

		{
			para869{1},
			ans869{true},
		},

		{
			para869{10},
			ans869{false},
		},

		{
			para869{16},
			ans869{true},
		},

		{
			para869{24},
			ans869{false},
		},

		{
			para869{46},
			ans869{true},
		},

		{
			para869{100},
			ans869{false},
		},

		{
			para869{123453242},
			ans869{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 869------------------------\n")

	for _, q := range qs {
		_, p := q.ans869, q.para869
		fmt.Printf("【input】:%v       【output】:%v\n", p, reorderedPowerOf2(p.n))
	}
	fmt.Printf("\n\n\n")
}
