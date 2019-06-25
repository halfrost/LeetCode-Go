package leetcode

import (
	"fmt"
	"testing"
)

type question433 struct {
	para433
	ans433
}

// para 是参数
// one 代表第一个参数
type para433 struct {
	start string
	end   string
	bank  []string
}

// ans 是答案
// one 代表第一个答案
type ans433 struct {
	one int
}

func Test_Problem433(t *testing.T) {

	qs := []question433{
		question433{
			para433{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}},
			ans433{1},
		},

		question433{
			para433{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}},
			ans433{2},
		},

		question433{
			para433{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}},
			ans433{3},
		},

		question433{
			para433{"AACCGGTT", "AACCGGTA", []string{}},
			ans433{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 433------------------------\n")

	for _, q := range qs {
		_, p := q.ans433, q.para433
		fmt.Printf("【input】:%v       【output】:%v\n", p, minMutation(p.start, p.end, p.bank))
	}
	fmt.Printf("\n\n\n")
}
