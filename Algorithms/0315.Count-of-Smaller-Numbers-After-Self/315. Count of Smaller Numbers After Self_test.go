package leetcode

import (
	"fmt"
	"testing"
)

type question315 struct {
	para315
	ans315
}

// para 是参数
// one 代表第一个参数
type para315 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans315 struct {
	one []int
}

func Test_Problem315(t *testing.T) {

	qs := []question315{

		question315{
			para315{[]int{5, 2, 6, 1}},
			ans315{[]int{2, 1, 1, 0}},
		},

		question315{
			para315{[]int{-1, -1}},
			ans315{[]int{0, 0}},
		},

		question315{
			para315{[]int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41}},
			ans315{[]int{10, 27, 10, 35, 12, 22, 28, 8, 19, 2, 12, 2, 9, 6, 12, 5, 17, 9, 19, 12, 14, 6, 12, 5, 12, 3, 0, 10, 0, 7, 8, 4, 0, 0, 4, 3, 2, 0, 1, 0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 315------------------------\n")

	for _, q := range qs {
		_, p := q.ans315, q.para315
		fmt.Printf("【input】:%v       【output】:%v\n", p, countSmaller(p.nums))
	}
	fmt.Printf("\n\n\n")
}
