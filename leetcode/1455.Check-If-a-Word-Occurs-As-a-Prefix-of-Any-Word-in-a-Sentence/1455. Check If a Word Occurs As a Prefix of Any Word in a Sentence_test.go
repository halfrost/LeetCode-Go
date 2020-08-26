package leetcode

import (
	"fmt"
	"testing"
)

type question1455 struct {
	para1455
	ans1455
}

// para 是参数
// one 代表第一个参数
type para1455 struct {
	sentence   string
	searchWord string
}

// ans 是答案
// one 代表第一个答案
type ans1455 struct {
	one int
}

func Test_Problem1455(t *testing.T) {

	qs := []question1455{

		{
			para1455{"i love eating burger", "burg"},
			ans1455{4},
		},

		{
			para1455{"this problem is an easy problem", "pro"},
			ans1455{2},
		},

		{
			para1455{"i am tired", "you"},
			ans1455{-1},
		},

		{
			para1455{"i use triple pillow", "pill"},
			ans1455{4},
		},

		{
			para1455{"hello from the other side", "they"},
			ans1455{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1455------------------------\n")

	for _, q := range qs {
		_, p := q.ans1455, q.para1455
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", isPrefixOfWord(p.sentence, p.searchWord))
	}
	fmt.Printf("\n\n\n")
}
