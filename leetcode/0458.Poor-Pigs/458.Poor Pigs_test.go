package leetcode

import (
	"fmt"
	"testing"
)

type question458 struct {
	para458
	ans458
}

// para 是参数
type para458 struct {
	buckets       int
	minutesToDie  int
	minutesToTest int
}

// ans 是答案
type ans458 struct {
	ans int
}

func Test_Problem458(t *testing.T) {

	qs := []question458{

		{
			para458{1000, 15, 60},
			ans458{5},
		},

		{
			para458{4, 15, 15},
			ans458{2},
		},

		{
			para458{4, 15, 30},
			ans458{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 458------------------------\n")

	for _, q := range qs {
		_, p := q.ans458, q.para458
		fmt.Printf("【input】:%v    【output】:%v\n", p, poorPigs(p.buckets, p.minutesToDie, p.minutesToTest))
	}
	fmt.Printf("\n\n\n")
}
