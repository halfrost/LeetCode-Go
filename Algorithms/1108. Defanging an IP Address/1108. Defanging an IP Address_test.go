package leetcode

import (
	"fmt"
	"testing"
)

type question1108 struct {
	para1108
	ans1108
}

// para 是参数
// one 代表第一个参数
type para1108 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1108 struct {
	one string
}

func Test_Problem1108(t *testing.T) {

	qs := []question1108{

		question1108{
			para1108{"1.1.1.1"},
			ans1108{"1[.]1[.]1[.]1"},
		},

		question1108{
			para1108{"255.100.50.0"},
			ans1108{"255[.]100[.]50[.]0"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1108------------------------\n")

	for _, q := range qs {
		_, p := q.ans1108, q.para1108
		fmt.Printf("【input】:%v       【output】:%v\n", p, defangIPaddr(p.one))
	}
	fmt.Printf("\n\n\n")
}
