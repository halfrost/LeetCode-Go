package leetcode

import (
	"fmt"
	"testing"
)

type question1078 struct {
	para1078
	ans1078
}

// para 是参数
// one 代表第一个参数
type para1078 struct {
	t string
	f string
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1078 struct {
	one []string
}

func Test_Problem1078(t *testing.T) {

	qs := []question1078{

		question1078{
			para1078{"alice is a good girl she is a good student", "a", "good"},
			ans1078{[]string{"girl", "student"}},
		},

		question1078{
			para1078{"we will we will rock you", "we", "will"},
			ans1078{[]string{"we", "rock"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1078------------------------\n")

	for _, q := range qs {
		_, p := q.ans1078, q.para1078
		fmt.Printf("【input】:%v       【output】:%v\n", p, findOcurrences(p.t, p.f, p.s))
	}
	fmt.Printf("\n\n\n")
}
