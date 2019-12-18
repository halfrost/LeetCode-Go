package leetcode

import (
	"fmt"
	"testing"
)

type question979 struct {
	para979
	ans979
}

// para 是参数
// one 代表第一个参数
type para979 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans979 struct {
	one int
}

func Test_Problem979(t *testing.T) {

	qs := []question979{

		question979{
			para979{[]int{}},
			ans979{0},
		},

		question979{
			para979{[]int{1}},
			ans979{0},
		},

		question979{
			para979{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			ans979{41},
		},

		question979{
			para979{[]int{1, 2, 3, 4, NULL, NULL, 5}},
			ans979{11},
		},

		question979{
			para979{[]int{1, 2, 3, 4, NULL, 5}},
			ans979{11},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 979------------------------\n")

	for _, q := range qs {
		_, p := q.ans979, q.para979
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", distributeCoins(root))
	}
	fmt.Printf("\n\n\n")
}
