package leetcode

import (
	"fmt"
	"testing"
)

type question48 struct {
	para48
	ans48
}

// para 是参数
// one 代表第一个参数
type para48 struct {
	s [][]int
}

// ans 是答案
// one 代表第一个答案
type ans48 struct {
	s [][]int
}

func Test_Problem48(t *testing.T) {

	qs := []question48{

		question48{
			para48{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}},
			ans48{[][]int{[]int{7, 4, 1}, []int{8, 5, 2}, []int{9, 6, 3}}},
		},

		question48{
			para48{[][]int{[]int{5, 1, 9, 11}, []int{2, 4, 8, 10}, []int{13, 3, 6, 7}, []int{15, 14, 12, 16}}},
			ans48{[][]int{[]int{15, 13, 2, 5}, []int{14, 3, 4, 1}, []int{12, 6, 8, 9}, []int{16, 7, 10, 11}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 48------------------------\n")

	for _, q := range qs {
		_, p := q.ans48, q.para48
		fmt.Printf("【input】:%v \n", p)
		rotate(p.s)
		fmt.Printf("【output】:%v\n", p)
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\n")
}
