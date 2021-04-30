package leetcode

import (
	"fmt"
	"testing"
)

type question690 struct {
	para690
	ans690
}

// para 是参数
// one 代表第一个参数
type para690 struct {
	employees []*Employee
	id        int
}

// ans 是答案
// one 代表第一个答案
type ans690 struct {
	one int
}

func Test_Problem690(t *testing.T) {

	qs := []question690{

		{
			para690{[]*Employee{{1, 5, []int{2, 3}}, {2, 3, []int{}}, {3, 3, []int{}}}, 1},
			ans690{11},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 690------------------------\n")

	for _, q := range qs {
		_, p := q.ans690, q.para690
		fmt.Printf("【input】:%v       【output】:%v\n", p, getImportance(p.employees, p.id))
	}
	fmt.Printf("\n\n\n")
}
