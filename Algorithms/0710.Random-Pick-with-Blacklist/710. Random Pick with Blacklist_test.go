package leetcode

import (
	"fmt"
	"testing"
)

type question710 struct {
	para710
	ans710
}

// para 是参数
// one 代表第一个参数
type para710 struct {
	n         int
	blackList []int
	times     int
}

// ans 是答案
// one 代表第一个答案
type ans710 struct {
	one []int
}

func Test_Problem710(t *testing.T) {

	qs := []question710{

		question710{
			para710{1, []int{}, 3},
			ans710{[]int{0, 0, 0}},
		},

		question710{
			para710{2, []int{}, 3},
			ans710{[]int{1, 1, 1}},
		},

		question710{
			para710{3, []int{1}, 3},
			ans710{[]int{0, 0, 2}},
		},

		question710{
			para710{4, []int{2}, 3},
			ans710{[]int{1, 3, 1}},
		},

		question710{
			para710{10000000, []int{1, 9999, 999999, 99999, 100, 0}, 10},
			ans710{[]int{400, 200, 300}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 710------------------------\n")

	for _, q := range qs {
		_, p := q.ans710, q.para710
		fmt.Printf("【input】: n = %v blacklist = %v pick times = %v  ", p.n, p.blackList, p.times)
		obj := Constructor710(p.n, p.blackList)
		fmt.Printf("【output】:")
		for i := 0; i < p.times; i++ {
			fmt.Printf(" %v ,", obj.Pick())
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\n")
}
