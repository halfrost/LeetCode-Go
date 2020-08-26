package leetcode

import (
	"fmt"
	"testing"
)

type question537 struct {
	para537
	ans537
}

// para 是参数
// one 代表第一个参数
type para537 struct {
	a string
	b string
}

// ans 是答案
// one 代表第一个答案
type ans537 struct {
	one string
}

func Test_Problem537(t *testing.T) {

	qs := []question537{

		{
			para537{"1+1i", "1+1i"},
			ans537{"0+2i"},
		},

		{
			para537{"1+-1i", "1+-1i"},
			ans537{"0+-2i"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 537------------------------\n")

	for _, q := range qs {
		_, p := q.ans537, q.para537
		fmt.Printf("【input】:%v       【output】:%v\n", p, complexNumberMultiply(p.a, p.b))
	}
	fmt.Printf("\n\n\n")
}
