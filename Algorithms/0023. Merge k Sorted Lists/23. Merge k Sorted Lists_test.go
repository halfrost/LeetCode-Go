package leetcode

import (
	"fmt"
	"testing"
)

type question23 struct {
	para23
	ans23
}

// para 是参数
// one 代表第一个参数
type para23 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans23 struct {
	one []int
}

func Test_Problem23(t *testing.T) {

	qs := []question23{

		question23{
			para23{[][]int{}},
			ans23{[]int{}},
		},

		question23{
			para23{[][]int{
				[]int{1},
				[]int{1},
			}},
			ans23{[]int{1, 1}},
		},

		question23{
			para23{[][]int{
				[]int{1, 2, 3, 4},
				[]int{1, 2, 3, 4},
			}},
			ans23{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		question23{
			para23{[][]int{
				[]int{1, 2, 3, 4, 5},
				[]int{1, 2, 3, 4, 5},
			}},
			ans23{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		question23{
			para23{[][]int{
				[]int{1},
				[]int{9, 9, 9, 9, 9},
			}},
			ans23{[]int{1, 9, 9, 9, 9, 9}},
		},

		question23{
			para23{[][]int{
				[]int{9, 9, 9, 9, 9},
				[]int{1},
			}},
			ans23{[]int{1, 9, 9, 9, 9, 9}},
		},

		question23{
			para23{[][]int{
				[]int{2, 3, 4},
				[]int{4, 5, 6},
			}},
			ans23{[]int{2, 3, 4, 4, 5, 6}},
		},

		question23{
			para23{[][]int{
				[]int{1, 3, 8},
				[]int{1, 7},
			}},
			ans23{[]int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 23------------------------\n")

	for _, q := range qs {
		var ls []*ListNode
		for _, qq := range q.para23.one {
			ls = append(ls, S2l(qq))
		}
		fmt.Printf("【input】:%v       【output】:%v\n", q.para23.one, L2s(mergeKLists(ls)))
	}
	fmt.Printf("\n\n\n")
}
