package leetcode

import (
	"fmt"
	"testing"
)

type question567 struct {
	para567
	ans567
}

// para 是参数
// one 代表第一个参数
type para567 struct {
	s string
	p string
}

// ans 是答案
// one 代表第一个答案
type ans567 struct {
	one bool
}

func Test_Problem567(t *testing.T) {

	qs := []question567{

		question567{
			para567{"ab", "abab"},
			ans567{true},
		},

		question567{
			para567{"abc", "cbaebabacd"},
			ans567{true},
		},

		question567{
			para567{"abc", ""},
			ans567{false},
		},

		question567{
			para567{"abc", "abacbabc"},
			ans567{true},
		},

		question567{
			para567{"ab", "eidboaoo"},
			ans567{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 567------------------------\n")

	for _, q := range qs {
		_, p := q.ans567, q.para567
		fmt.Printf("【input】:%v       【output】:%v\n", p, checkInclusion(p.s, p.p))
	}
	fmt.Printf("\n\n\n")
}
