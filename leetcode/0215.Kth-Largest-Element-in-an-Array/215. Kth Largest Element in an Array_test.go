package leetcode

import (
	"fmt"
	"testing"
)

type question215 struct {
	para215
	ans215
}

// para 是参数
// one 代表第一个参数
type para215 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans215 struct {
	one int
}

func Test_Problem215(t *testing.T) {

	qs := []question215{

		question215{
			para215{[]int{3, 2, 1, 5, 6, 4}, 2},
			ans215{5},
		},

		question215{
			para215{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
			ans215{4},
		},

		question215{
			para215{[]int{0, 0, 0, 0, 0}, 2},
			ans215{0},
		},

		question215{
			para215{[]int{1}, 1},
			ans215{1},
		},

		question215{
			para215{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 20},
			ans215{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 215------------------------\n")

	for _, q := range qs {
		_, p := q.ans215, q.para215
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, findKthLargest(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
