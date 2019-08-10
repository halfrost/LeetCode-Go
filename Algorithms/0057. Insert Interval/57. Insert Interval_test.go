package leetcode

import (
	"fmt"
	"testing"
)

type question57 struct {
	para57
	ans57
}

// para 是参数
// one 代表第一个参数
type para57 struct {
	one []Interval
	two Interval
}

// ans 是答案
// one 代表第一个答案
type ans57 struct {
	one []Interval
}

func Test_Problem57(t *testing.T) {

	qs := []question57{

		question57{
			para57{[]Interval{}, Interval{}},
			ans57{[]Interval{}},
		},

		question57{
			para57{[]Interval{Interval{Start: 1, End: 3}, Interval{Start: 6, End: 9}}, Interval{Start: 4, End: 8}},
			ans57{[]Interval{Interval{Start: 1, End: 5}, Interval{Start: 6, End: 9}}},
		},

		question57{
			para57{[]Interval{Interval{Start: 1, End: 3}, Interval{Start: 6, End: 9}}, Interval{Start: 2, End: 5}},
			ans57{[]Interval{Interval{Start: 1, End: 5}, Interval{Start: 6, End: 9}}},
		},

		question57{
			para57{[]Interval{Interval{Start: 1, End: 2}, Interval{Start: 3, End: 5}, Interval{Start: 6, End: 7}, Interval{Start: 8, End: 10}, Interval{Start: 12, End: 16}}, Interval{Start: 4, End: 8}},
			ans57{[]Interval{Interval{Start: 1, End: 2}, Interval{Start: 3, End: 10}, Interval{Start: 12, End: 16}}},
		},

		question57{
			para57{[]Interval{Interval{Start: 1, End: 5}}, Interval{Start: 5, End: 7}},
			ans57{[]Interval{Interval{Start: 1, End: 7}}},
		},

		question57{
			para57{[]Interval{Interval{Start: 1, End: 2}, Interval{Start: 3, End: 5}, Interval{Start: 6, End: 7}, Interval{Start: 8, End: 10}, Interval{Start: 12, End: 16}}, Interval{Start: 9, End: 12}},
			ans57{[]Interval{Interval{Start: 1, End: 2}, Interval{Start: 3, End: 5}, Interval{Start: 6, End: 7}, Interval{Start: 8, End: 16}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 57------------------------\n")

	for _, q := range qs {
		_, p := q.ans57, q.para57
		fmt.Printf("【input】:%v       【output】:%v\n", p, insert(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
