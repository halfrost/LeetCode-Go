package leetcode

import (
	"fmt"
	"testing"
)

type question513 struct {
	para513
	ans513
}

// para 是参数
// one 代表第一个参数
type para513 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans513 struct {
	one int
}

func Test_Problem513(t *testing.T) {

	qs := []question513{

		question513{
			para513{[]int{}},
			ans513{0},
		},

		question513{
			para513{[]int{1}},
			ans513{1},
		},

		question513{
			para513{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			ans513{15},
		},

		question513{
			para513{[]int{1, 2, 3, 4, NULL, NULL, 5}},
			ans513{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 513------------------------\n")

	for _, q := range qs {
		_, p := q.ans513, q.para513
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", findBottomLeftValue(root))
	}
	fmt.Printf("\n\n\n")
}
