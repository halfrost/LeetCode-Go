package leetcode

import (
	"fmt"
	"testing"
)

type question767 struct {
	para767
	ans767
}

// para 是参数
// one 代表第一个参数
type para767 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans767 struct {
	one string
}

func Test_Problem767(t *testing.T) {

	qs := []question767{
		question767{
			para767{"aab"},
			ans767{"aba"},
		},

		question767{
			para767{"aaab"},
			ans767{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 767------------------------\n")

	for _, q := range qs {
		_, p := q.ans767, q.para767
		fmt.Printf("【input】:%v       【output】:%v\n", p, reorganizeString(p.one))
	}
	fmt.Printf("\n\n\n")
}
