package leetcode

import (
	"fmt"
	"testing"
)

type question990 struct {
	para990
	ans990
}

// para 是参数
// one 代表第一个参数
type para990 struct {
	a []string
}

// ans 是答案
// one 代表第一个答案
type ans990 struct {
	one bool
}

func Test_Problem990(t *testing.T) {

	qs := []question990{

		question990{
			para990{[]string{"a==b", "b!=a"}},
			ans990{false},
		},

		question990{
			para990{[]string{"b==a", "a==b"}},
			ans990{true},
		},

		question990{
			para990{[]string{"a==b", "b==c", "a==c"}},
			ans990{true},
		},

		question990{
			para990{[]string{"a==b", "b!=c", "c==a"}},
			ans990{false},
		},

		question990{
			para990{[]string{"c==c", "b==d", "x!=z"}},
			ans990{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 990------------------------\n")

	for _, q := range qs {
		_, p := q.ans990, q.para990
		fmt.Printf("【input】:%v       【output】:%v\n", p, equationsPossible(p.a))
	}
	fmt.Printf("\n\n\n")
}
