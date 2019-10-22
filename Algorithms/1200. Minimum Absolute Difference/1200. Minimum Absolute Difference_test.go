package leetcode

import (
	"fmt"
	"testing"
)

type question1200 struct {
	para1200
	ans1200
}

// para 是参数
// one 代表第一个参数
type para1200 struct {
	arr []int
}

// ans 是答案
// one 代表第一个答案
type ans1200 struct {
	one [][]int
}

func Test_Problem1200(t *testing.T) {

	qs := []question1200{

		question1200{
			para1200{[]int{4, 2, 1, 3}},
			ans1200{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}}},
		},

		question1200{
			para1200{[]int{1, 3, 6, 10, 15}},
			ans1200{[][]int{[]int{1, 3}}},
		},

		question1200{
			para1200{[]int{3, 8, -10, 23, 19, -4, -14, 27}},
			ans1200{[][]int{[]int{-14, -10}, []int{19, 23}, []int{23, 27}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1200------------------------\n")

	for _, q := range qs {
		_, p := q.ans1200, q.para1200
		fmt.Printf("【input】:%v       【output】:%v\n", p, minimumAbsDifference(p.arr))
	}
	fmt.Printf("\n\n\n")
}
