package leetcode

import (
	"fmt"
	"testing"
)

type question638 struct {
	para638
	ans638
}

// para 是参数
// one 代表第一个参数
type para638 struct {
	price   []int
	special [][]int
	needs   []int
}

// ans 是答案
// one 代表第一个答案
type ans638 struct {
	one int
}

func Test_Problem638(t *testing.T) {

	qs := []question638{

		{
			para638{[]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}},
			ans638{14},
		},

		{
			para638{[]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}},
			ans638{11},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 638------------------------\n")

	for _, q := range qs {
		_, p := q.ans638, q.para638
		fmt.Printf("【input】:%v       【output】:%v\n", p, shoppingOffers(p.price, p.special, p.needs))
	}
	fmt.Printf("\n\n\n")
}
