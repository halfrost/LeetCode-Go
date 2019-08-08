package leetcode

import (
	"fmt"
	"testing"
)

type question572 struct {
	para572
	ans572
}

// para 是参数
// one 代表第一个参数
type para572 struct {
	s []int
	t []int
}

// ans 是答案
// one 代表第一个答案
type ans572 struct {
	one bool
}

func Test_Problem572(t *testing.T) {

	qs := []question572{

		question572{
			para572{[]int{}, []int{}},
			ans572{false},
		},

		question572{
			para572{[]int{3, 4, 5, 1, 2}, []int{4, 1, 2}},
			ans572{true},
		},

		question572{
			para572{[]int{1, 1}, []int{1}},
			ans572{true},
		},

		question572{
			para572{[]int{1, NULL, 1, NULL, 1, NULL, 1, NULL, 1, NULL, 1, NULL, 1, NULL, 1, NULL, 1, NULL, 1, NULL, 1, 2}, []int{1, NULL, 1, NULL, 1, NULL, 1, NULL, 1, NULL, 1, 2}},
			ans572{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 572------------------------\n")

	for _, q := range qs {
		_, p := q.ans572, q.para572
		fmt.Printf("【input】:%v      ", p)
		roots := Ints2TreeNode(p.s)
		roott := Ints2TreeNode(p.t)
		fmt.Printf("【output】:%v      \n", isSubtree(roots, roott))
	}
	fmt.Printf("\n\n\n")
}
