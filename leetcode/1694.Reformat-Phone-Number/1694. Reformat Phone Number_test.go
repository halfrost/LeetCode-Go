package leetcode

import (
	"fmt"
	"testing"
)

type question1694 struct {
	para1694
	ans1694
}

// para 是参数
// one 代表第一个参数
type para1694 struct {
	number string
}

// ans 是答案
// one 代表第一个答案
type ans1694 struct {
	one string
}

func Test_Problem1694(t *testing.T) {

	qs := []question1694{

		{
			para1694{"1-23-45 6"},
			ans1694{"123-456"},
		},

		{
			para1694{"123 4-567"},
			ans1694{"123-45-67"},
		},

		{
			para1694{"123 4-5678"},
			ans1694{"123-456-78"},
		},

		{
			para1694{"12"},
			ans1694{"12"},
		},

		{
			para1694{"--17-5 229 35-39475 "},
			ans1694{"175-229-353-94-75"},
		},

		{
			para1694{"9964-"},
			ans1694{"99-64"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1694------------------------\n")

	for _, q := range qs {
		_, p := q.ans1694, q.para1694
		fmt.Printf("【input】:%v       【output】:%v\n", p, reformatNumber(p.number))
	}
	fmt.Printf("\n\n\n")
}
