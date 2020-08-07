package leetcode

import (
	"fmt"
	"testing"
)

type question402 struct {
	para402
	ans402
}

// para 是参数
// one 代表第一个参数
type para402 struct {
	num string
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans402 struct {
	one string
}

func Test_Problem402(t *testing.T) {

	qs := []question402{
		question402{
			para402{"10", 1},
			ans402{"0"},
		},

		question402{
			para402{"1111111", 3},
			ans402{"1111"},
		},

		question402{
			para402{"5337", 2},
			ans402{"33"},
		},

		question402{
			para402{"112", 1},
			ans402{"11"},
		},

		question402{
			para402{"1432219", 3},
			ans402{"1219"},
		},

		question402{
			para402{"10200", 1},
			ans402{"200"},
		},

		question402{
			para402{"10", 2},
			ans402{"0"},
		},

		question402{
			para402{"19", 2},
			ans402{"0"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 402------------------------\n")

	for _, q := range qs {
		_, p := q.ans402, q.para402
		fmt.Printf("【input】:%v       【output】:%v\n", p, removeKdigits(p.num, p.k))
	}
	fmt.Printf("\n\n\n")
}
