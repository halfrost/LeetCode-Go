package leetcode

import (
	"fmt"
	"testing"
)

type question179 struct {
	para179
	ans179
}

// para 是参数
// one 代表第一个参数
type para179 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans179 struct {
	one string
}

func Test_Problem179(t *testing.T) {

	qs := []question179{

		question179{
			para179{[]int{3, 6, 9, 1}},
			ans179{"9631"},
		},
		question179{
			para179{[]int{1}},
			ans179{"1"},
		},

		question179{
			para179{[]int{}},
			ans179{""},
		},

		question179{
			para179{[]int{2, 10}},
			ans179{"210"},
		},

		question179{
			para179{[]int{3, 30, 34, 5, 9}},
			ans179{"9534330"},
		},

		question179{
			para179{[]int{12, 128}},
			ans179{"12812"},
		},

		question179{
			para179{[]int{12, 121}},
			ans179{"12121"},
		},

		question179{
			para179{[]int{0, 0}},
			ans179{"0"},
		},

		question179{
			para179{[]int{1440, 7548, 4240, 6616, 733, 4712, 883, 8, 9576}},
			ans179{"9576888375487336616471242401440"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 179------------------------\n")

	for _, q := range qs {
		_, p := q.ans179, q.para179
		fmt.Printf("【input】:%v       【output】:%v\n", p, largestNumber(p.one))
	}
	fmt.Printf("\n\n\n")
}
