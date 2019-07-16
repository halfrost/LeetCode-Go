package leetcode

import (
	"fmt"
	"testing"
)

type question771 struct {
	para771
	ans771
}

// para 是参数
// one 代表第一个参数
type para771 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans771 struct {
	one int
}

func Test_Problem771(t *testing.T) {

	qs := []question771{
		question771{
			para771{"aA", "aAAbbbb"},
			ans771{3},
		},

		question771{
			para771{"z", "ZZ"},
			ans771{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 771------------------------\n")

	for _, q := range qs {
		_, p := q.ans771, q.para771
		fmt.Printf("【input】:%v       【output】:%v\n", p, numJewelsInStones(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
