package leetcode

import (
	"fmt"
	"testing"
)

type question532 struct {
	para532
	ans532
}

// para 是参数
// one 代表第一个参数
type para532 struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans532 struct {
	one int
}

func Test_Problem532(t *testing.T) {

	qs := []question532{

		question532{
			para532{[]int{3, 1, 4, 1, 5}, 2},
			ans532{2},
		},

		question532{
			para532{[]int{1, 2, 3, 4, 5}, 1},
			ans532{4},
		},

		question532{
			para532{[]int{1, 3, 1, 5, 4}, 0},
			ans532{1},
		},

		question532{
			para532{[]int{}, 3},
			ans532{0},
		},

		// question532{
		// 	para532{[]int{1, 2, 3, 2, 3, 2, 3, 2}, 0},
		// 	ans532{[]int{1, 2, 3, 2, 3, 2, 3, 2}},
		// },

		// question532{
		// 	para532{[]int{1, 2, 3, 4, 5}, 5},
		// 	ans532{[]int{1, 2, 3, 4}},
		// },

		// question532{
		// 	para532{[]int{}, 5},
		// 	ans532{[]int{}},
		// },

		// question532{
		// 	para532{[]int{1, 2, 3, 4, 5}, 10},
		// 	ans532{[]int{1, 2, 3, 4, 5}},
		// },

		// question532{
		// 	para532{[]int{1}, 1},
		// 	ans532{[]int{}},
		// },
	}

	fmt.Printf("------------------------Leetcode Problem 532------------------------\n")

	for _, q := range qs {
		_, p := q.ans532, q.para532
		fmt.Printf("【input】:%v       【output】:%v\n", p, findPairs(p.one, p.n))
	}
	fmt.Printf("\n\n\n")
}
