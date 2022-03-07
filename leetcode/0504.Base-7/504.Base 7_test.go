package leetcode

import (
	"fmt"
	"testing"
)

type question504 struct {
	para504
	ans504
}

// para 是参数
type para504 struct {
	num int
}

// ans 是答案
type ans504 struct {
	ans string
}

func Test_Problem504(t *testing.T) {

	qs := []question504{

		{
			para504{100},
			ans504{"202"},
		},

		{
			para504{-7},
			ans504{"-10"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 504------------------------\n")

	for _, q := range qs {
		_, p := q.ans504, q.para504
		fmt.Printf("【input】:%v      ", p.num)
		fmt.Printf("【output】:%v      \n", convertToBase7(p.num))
	}
	fmt.Printf("\n\n\n")
}
