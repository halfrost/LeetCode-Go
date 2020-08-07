package leetcode

import (
	"fmt"
	"testing"
)

type question1079 struct {
	para1079
	ans1079
}

// para 是参数
// one 代表第一个参数
type para1079 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1079 struct {
	one int
}

func Test_Problem1079(t *testing.T) {

	qs := []question1079{

		question1079{
			para1079{"AAB"},
			ans1079{8},
		},

		question1079{
			para1079{"AAABBC"},
			ans1079{188},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1079------------------------\n")

	for _, q := range qs {
		_, p := q.ans1079, q.para1079
		fmt.Printf("【input】:%v       【output】:%v\n", p, numTilePossibilities(p.one))
	}
	fmt.Printf("\n\n\n")
}
