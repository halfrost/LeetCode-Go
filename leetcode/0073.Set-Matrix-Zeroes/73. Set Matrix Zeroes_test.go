package leetcode

import (
	"fmt"
	"testing"
)

type question73 struct {
	para73
	ans73
}

// para 是参数
// one 代表第一个参数
type para73 struct {
	matrix [][]int
}

// ans 是答案
// one 代表第一个答案
type ans73 struct {
}

func Test_Problem73(t *testing.T) {

	qs := []question73{

		{
			para73{[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			}},
			ans73{},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 73------------------------\n")

	for _, q := range qs {
		_, p := q.ans73, q.para73
		fmt.Printf("【input】:%v       ", p)
		setZeroes(p.matrix)
		fmt.Printf("【output】:%v\n", p)
	}
	fmt.Printf("\n\n\n")
}
