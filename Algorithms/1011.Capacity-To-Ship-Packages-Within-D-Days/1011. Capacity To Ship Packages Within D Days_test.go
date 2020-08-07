package leetcode

import (
	"fmt"
	"testing"
)

type question1011 struct {
	para1011
	ans1011
}

// para 是参数
// one 代表第一个参数
type para1011 struct {
	weights []int
	D       int
}

// ans 是答案
// one 代表第一个答案
type ans1011 struct {
	one int
}

func Test_Problem1011(t *testing.T) {

	qs := []question1011{

		question1011{
			para1011{[]int{7, 2, 5, 10, 8}, 2},
			ans1011{18},
		},

		question1011{
			para1011{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
			ans1011{15},
		},

		question1011{
			para1011{[]int{3, 2, 2, 4, 1, 4}, 3},
			ans1011{6},
		},

		question1011{
			para1011{[]int{1, 2, 3, 1, 1}, 4},
			ans1011{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1011------------------------\n")

	for _, q := range qs {
		_, p := q.ans1011, q.para1011
		fmt.Printf("【input】:%v       【output】:%v\n", p, shipWithinDays(p.weights, p.D))
	}
	fmt.Printf("\n\n\n")
}
