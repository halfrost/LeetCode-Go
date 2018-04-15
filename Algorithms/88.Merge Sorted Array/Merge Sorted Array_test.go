package problem412

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
	m   int
	two []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0088(t *testing.T) {

	qs := []question{

		// question{
		// 	para{[]int{0}, 0, []int{1}, 1},
		// 	ans{[]int{1}},
		// },
		//
		// question{
		// 	para{[]int{1, 3, 5, 7}, 4, []int{2, 4}, 2},
		// 	ans{[]int{1, 2, 3, 4}},
		// },
		//
		// question{
		// 	para{[]int{1, 3, 5, 7}, 4, []int{2, 2}, 2},
		// 	ans{[]int{1, 2, 2, 3}},
		// },

		question{
			para{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			ans{[]int{1, 2, 2, 3}},
		},
	}

	for _, q := range qs {
		_, p := q.ans, q.para

		merge(p.one, p.m, p.two, p.n)
		fmt.Printf("~~%v~~\n", p)
	}
}
