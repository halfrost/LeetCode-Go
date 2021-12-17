package leetcode

import (
	"fmt"
	"testing"
)

type question1518 struct {
	para1518
	ans1518
}

// para 是参数
type para1518 struct {
	numBottles  int
	numExchange int
}

// ans 是答案
type ans1518 struct {
	ans int
}

func Test_Problem1518(t *testing.T) {

	qs := []question1518{

		{
			para1518{9, 3},
			ans1518{13},
		},

		{
			para1518{15, 4},
			ans1518{19},
		},

		{
			para1518{5, 5},
			ans1518{6},
		},

		{
			para1518{2, 3},
			ans1518{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1518------------------------\n")

	for _, q := range qs {
		_, p := q.ans1518, q.para1518
		fmt.Printf("【input】:%v    【output】:%v\n", p, numWaterBottles(p.numBottles, p.numExchange))
	}
	fmt.Printf("\n\n\n")
}
