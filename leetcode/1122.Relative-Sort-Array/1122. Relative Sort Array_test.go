package leetcode

import (
	"fmt"
	"testing"
)

type question1122 struct {
	para1122
	ans1122
}

// para 是参数
// one 代表第一个参数
type para1122 struct {
	arr1 []int
	arr2 []int
}

// ans 是答案
// one 代表第一个答案
type ans1122 struct {
	one []int
}

func Test_Problem1122(t *testing.T) {

	qs := []question1122{

		question1122{
			para1122{[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}},
			ans1122{[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1122------------------------\n")

	for _, q := range qs {
		_, p := q.ans1122, q.para1122
		fmt.Printf("【input】:%v       【output】:%v\n", p, relativeSortArray(p.arr1, p.arr2))
	}
	fmt.Printf("\n\n\n")
}
