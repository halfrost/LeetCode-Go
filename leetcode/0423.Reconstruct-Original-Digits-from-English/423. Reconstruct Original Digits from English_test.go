package leetcode

import (
	"fmt"
	"testing"
)

type question423 struct {
	para423
	ans423
}

// para 是参数
// one 代表第一个参数
type para423 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans423 struct {
	one int
}

func Test_Problem423(t *testing.T) {

	qs := []question423{

		{
			para423{"owoztneoer"},
			ans423{012},
		},

		{
			para423{"fviefuro"},
			ans423{45},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 423------------------------\n")

	for _, q := range qs {
		_, p := q.ans423, q.para423
		fmt.Printf("【input】:%v       【output】:%v\n", p, originalDigits(p.one))
	}
	fmt.Printf("\n\n\n")
}
