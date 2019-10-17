package leetcode

import (
	"fmt"
	"testing"
)

type question852 struct {
	para852
	ans852
}

// para 是参数
// one 代表第一个参数
type para852 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans852 struct {
	one int
}

func Test_Problem852(t *testing.T) {

	qs := []question852{

		question852{
			para852{[]int{0, 1, 0}},
			ans852{1},
		},

		question852{
			para852{[]int{0, 2, 1, 0}},
			ans852{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 852------------------------\n")

	for _, q := range qs {
		_, p := q.ans852, q.para852
		fmt.Printf("【input】:%v       【output】:%v\n", p, peakIndexInMountainArray(p.one))
	}
	fmt.Printf("\n\n\n")
}
