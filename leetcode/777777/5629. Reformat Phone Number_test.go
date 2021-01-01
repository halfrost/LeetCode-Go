package leetcode

import (
	"fmt"
	"testing"
)

type question1690 struct {
	para1690
	ans1690
}

// para 是参数
// one 代表第一个参数
type para1690 struct {
	number string
}

// ans 是答案
// one 代表第一个答案
type ans1690 struct {
	one string
}

func Test_Problem1690(t *testing.T) {

	qs := []question1690{

		{
			para1690{"1-23-45 6"},
			ans1690{"123-456"},
		},

		{
			para1690{"123 4-567"},
			ans1690{"123-45-67"},
		},

		{
			para1690{"123 4-5678"},
			ans1690{"123-456-78"},
		},

		{
			para1690{"12"},
			ans1690{"12"},
		},

		{
			para1690{"--17-5 229 35-39475 "},
			ans1690{"175-229-353-94-75"},
		},

		{
			para1690{"9964-"},
			ans1690{"99-64"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1690------------------------\n")

	for _, q := range qs {
		_, p := q.ans1690, q.para1690
		fmt.Printf("【input】:%v       【output】:%v\n", p, reformatNumber(p.number))
	}
	fmt.Printf("\n\n\n")
}
