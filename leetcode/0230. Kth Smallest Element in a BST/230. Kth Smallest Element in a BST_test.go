package leetcode

import (
	"fmt"
	"testing"
)

type question230 struct {
	para230
	ans230
}

// para 是参数
// one 代表第一个参数
type para230 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans230 struct {
	one int
}

func Test_Problem230(t *testing.T) {

	qs := []question230{

		question230{
			para230{[]int{}, 0},
			ans230{0},
		},

		question230{
			para230{[]int{3, 1, 4, NULL, 2}, 1},
			ans230{1},
		},

		question230{
			para230{[]int{5, 3, 6, 2, 4, NULL, NULL, 1}, 3},
			ans230{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 230------------------------\n")

	for _, q := range qs {
		_, p := q.ans230, q.para230
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", kthSmallest(root, p.k))
	}
	fmt.Printf("\n\n\n")
}
