package leetcode

import (
	"fmt"
	"testing"
)

type question85 struct {
	para85
	ans85
}

// para 是参数
// one 代表第一个参数
type para85 struct {
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans85 struct {
	one int
}

func Test_Problem85(t *testing.T) {

	qs := []question85{

		{
			para85{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}},
			ans85{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 85------------------------\n")

	for _, q := range qs {
		_, p := q.ans85, q.para85
		fmt.Printf("【input】:%v       【output】:%v\n", p, maximalRectangle(p.one))
	}
	fmt.Printf("\n\n\n")
}
