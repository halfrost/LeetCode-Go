package leetcode

import (
	"fmt"
	"testing"
)

type question108 struct {
	para108
	ans108
}

// para 是参数
// one 代表第一个参数
type para108 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans108 struct {
	one []int
}

func Test_Problem108(t *testing.T) {

	qs := []question108{

		question108{
			para108{[]int{-10, -3, 0, 5, 9}},
			ans108{[]int{0, -3, 9, -10, NULL, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 108------------------------\n")

	for _, q := range qs {
		_, p := q.ans108, q.para108
		arr := []int{}
		T2s(sortedArrayToBST(p.one), &arr)
		fmt.Printf("【input】:%v       【output】:%v\n", p, arr)
	}
	fmt.Printf("\n\n\n")
}
