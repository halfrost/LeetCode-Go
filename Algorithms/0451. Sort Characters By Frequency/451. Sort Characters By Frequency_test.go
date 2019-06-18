package leetcode

import (
	"fmt"
	"testing"
)

type question451 struct {
	para451
	ans451
}

// para 是参数
// one 代表第一个参数
type para451 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans451 struct {
	one string
}

func Test_Problem451(t *testing.T) {

	qs := []question451{
		question451{
			para451{"tree"},
			ans451{"eert"},
		},

		question451{
			para451{"cccaaa"},
			ans451{"cccaaa"},
		},

		question451{
			para451{"Aabb"},
			ans451{"bbAa"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 451------------------------\n")

	for _, q := range qs {
		_, p := q.ans451, q.para451
		fmt.Printf("【input】:%v       【output】:%v\n", p, frequencySort(p.one))
	}
	fmt.Printf("\n\n\n")
}
