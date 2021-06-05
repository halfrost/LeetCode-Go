package leetcode

import (
	"fmt"
	"testing"
)

type question1383 struct {
	para1383
	ans1383
}

// para 是参数
// one 代表第一个参数
type para1383 struct {
	n          int
	speed      []int
	efficiency []int
	k          int
}

// ans 是答案
// one 代表第一个答案
type ans1383 struct {
	one int
}

func Test_Problem1383(t *testing.T) {

	qs := []question1383{

		{
			para1383{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2},
			ans1383{60},
		},

		{
			para1383{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3},
			ans1383{68},
		},

		{
			para1383{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4},
			ans1383{72},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1383------------------------\n")

	for _, q := range qs {
		_, p := q.ans1383, q.para1383
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxPerformance(p.n, p.speed, p.efficiency, p.k))
	}
	fmt.Printf("\n\n\n")
}
