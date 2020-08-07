package leetcode

import (
	"fmt"
	"testing"
)

type question1184 struct {
	para1184
	ans1184
}

// para 是参数
// one 代表第一个参数
type para1184 struct {
	distance    []int
	start       int
	destination int
}

// ans 是答案
// one 代表第一个答案
type ans1184 struct {
	one int
}

func Test_Problem1184(t *testing.T) {

	qs := []question1184{

		question1184{
			para1184{[]int{1, 2, 3, 4}, 0, 1},
			ans1184{1},
		},

		question1184{
			para1184{[]int{1, 2, 3, 4}, 0, 2},
			ans1184{3},
		},

		question1184{
			para1184{[]int{1, 2, 3, 4}, 0, 3},
			ans1184{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1184------------------------\n")

	for _, q := range qs {
		_, p := q.ans1184, q.para1184
		fmt.Printf("【input】:%v       【output】:%v\n", p, distanceBetweenBusStops(p.distance, p.start, p.destination))
	}
	fmt.Printf("\n\n\n")
}
