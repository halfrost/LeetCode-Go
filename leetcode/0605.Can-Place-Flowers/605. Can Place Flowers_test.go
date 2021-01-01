package leetcode

import (
	"fmt"
	"testing"
)

type question605 struct {
	para605
	ans605
}

// para 是参数
// one 代表第一个参数
type para605 struct {
	flowerbed []int
	n         int
}

// ans 是答案
// one 代表第一个答案
type ans605 struct {
	one bool
}

func Test_Problem605(t *testing.T) {

	qs := []question605{

		{
			para605{[]int{1, 0, 0, 0, 1}, 1},
			ans605{true},
		},

		{
			para605{[]int{1, 0, 0, 0, 1}, 2},
			ans605{false},
		},

		{
			para605{[]int{1, 0, 0, 0, 0, 1}, 2},
			ans605{false},
		},

		{
			para605{[]int{0, 0, 1, 0}, 1},
			ans605{true},
		},

		{
			para605{[]int{0, 0, 1, 0, 0}, 1},
			ans605{true},
		},

		{
			para605{[]int{1, 0, 0, 1, 0}, 2},
			ans605{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 605------------------------\n")

	for _, q := range qs {
		_, p := q.ans605, q.para605
		fmt.Printf("【input】:%v       【output】:%v\n", p, canPlaceFlowers(p.flowerbed, p.n))
	}
	fmt.Printf("\n\n\n")
}
