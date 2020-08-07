package leetcode

import (
	"fmt"
	"testing"
)

type question881 struct {
	para881
	ans881
}

// para 是参数
// one 代表第一个参数
type para881 struct {
	s []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans881 struct {
	one int
}

func Test_Problem881(t *testing.T) {

	qs := []question881{

		question881{
			para881{[]int{1, 2}, 3},
			ans881{1},
		},

		question881{
			para881{[]int{3, 2, 2, 1}, 3},
			ans881{3},
		},

		question881{
			para881{[]int{3, 5, 3, 4}, 5},
			ans881{4},
		},

		question881{
			para881{[]int{5, 1, 4, 2}, 6},
			ans881{2},
		},

		question881{
			para881{[]int{3, 2, 2, 1}, 3},
			ans881{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 881------------------------\n")

	for _, q := range qs {
		_, p := q.ans881, q.para881
		fmt.Printf("【input】:%v       【output】:%v\n", p, numRescueBoats(p.s, p.k))
	}
	fmt.Printf("\n\n\n")
}
