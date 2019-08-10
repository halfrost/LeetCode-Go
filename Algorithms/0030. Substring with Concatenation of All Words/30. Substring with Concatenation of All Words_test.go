package leetcode

import (
	"fmt"
	"testing"
)

type question30 struct {
	para30
	ans30
}

// para 是参数
// one 代表第一个参数
type para30 struct {
	one string
	two []string
}

// ans 是答案
// one 代表第一个答案
type ans30 struct {
	one []int
}

func Test_Problem30(t *testing.T) {

	qs := []question30{

		question30{
			para30{"aaaaaaaa", []string{"aa", "aa", "aa"}},
			ans30{[]int{0, 1, 2}},
		},

		question30{
			para30{"barfoothefoobarman", []string{"foo", "bar"}},
			ans30{[]int{0, 9}},
		},

		question30{
			para30{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}},
			ans30{[]int{}},
		},

		question30{
			para30{"goodgoodgoodgoodgood", []string{"good"}},
			ans30{[]int{0, 4, 8, 12, 16}},
		},

		question30{
			para30{"barofoothefoolbarman", []string{"foo", "bar"}},
			ans30{[]int{}},
		},

		question30{
			para30{"bbarffoothefoobarman", []string{"foo", "bar"}},
			ans30{[]int{}},
		},

		question30{
			para30{"ooroodoofoodtoo", []string{"foo", "doo", "roo", "tee", "oo"}},
			ans30{[]int{}},
		},

		question30{
			para30{"abc", []string{"a", "b", "c"}},
			ans30{[]int{0}},
		},

		question30{
			para30{"a", []string{"b"}},
			ans30{[]int{}},
		},

		question30{
			para30{"ab", []string{"ba"}},
			ans30{[]int{}},
		},

		question30{
			para30{"n", []string{}},
			ans30{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 30------------------------\n")
	for _, q := range qs {
		_, p := q.ans30, q.para30
		fmt.Printf("【input】:%v       【output】:%v\n", p, findSubstring(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
