package leetcode

import (
	"fmt"
	"testing"
)

type question21 struct {
	para21
	ans21
}

// para 是参数
// one 代表第一个参数
type para21 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans21 struct {
	one []int
}

func Test_Problem21(t *testing.T) {

	qs := []question21{

		question21{
			para21{[]int{}, []int{}},
			ans21{[]int{}},
		},

		question21{
			para21{[]int{1}, []int{1}},
			ans21{[]int{1, 1}},
		},

		question21{
			para21{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans21{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		question21{
			para21{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans21{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		question21{
			para21{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans21{[]int{1, 9, 9, 9, 9, 9}},
		},

		question21{
			para21{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans21{[]int{1, 9, 9, 9, 9, 9}},
		},

		question21{
			para21{[]int{2, 3, 4}, []int{4, 5, 6}},
			ans21{[]int{2, 3, 4, 4, 5, 6}},
		},

		question21{
			para21{[]int{1, 3, 8}, []int{1, 7}},
			ans21{[]int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 21------------------------\n")

	for _, q := range qs {
		_, p := q.ans21, q.para21
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(mergeTwoLists(S2l(p.one), S2l(p.another))))
	}
	fmt.Printf("\n\n\n")
}
