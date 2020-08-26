package leetcode

import (
	"fmt"
	"testing"
)

type question1385 struct {
	para1385
	ans1385
}

// para 是参数
// one 代表第一个参数
type para1385 struct {
	arr1 []int
	arr2 []int
	d    int
}

// ans 是答案
// one 代表第一个答案
type ans1385 struct {
	one []int
}

func Test_Problem1385(t *testing.T) {

	qs := []question1385{

		{
			para1385{[]int{4, 5, 8}, []int{10, 9, 1, 8}, 2},
			ans1385{[]int{2}},
		},

		{
			para1385{[]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3},
			ans1385{[]int{2}},
		},

		{
			para1385{[]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6},
			ans1385{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1385------------------------\n")

	for _, q := range qs {
		_, p := q.ans1385, q.para1385
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", findTheDistanceValue(p.arr1, p.arr2, p.d))
	}
	fmt.Printf("\n\n\n")
}
