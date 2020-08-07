package leetcode

import (
	"fmt"
	"testing"
)

type question784 struct {
	para784
	ans784
}

// para 是参数
// one 代表第一个参数
type para784 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans784 struct {
	one []string
}

func Test_Problem784(t *testing.T) {

	qs := []question784{

		question784{
			para784{"mQe"},
			ans784{[]string{"mqe", "mqE", "mQe", "mQE", "Mqe", "MqE", "MQe", "MQE"}},
		},

		question784{
			para784{"C"},
			ans784{[]string{"c", "C"}},
		},

		question784{
			para784{"a1b2"},
			ans784{[]string{"a1b2", "a1B2", "A1b2", "A1B2"}},
		},

		question784{
			para784{"3z4"},
			ans784{[]string{"3z4", "3Z4"}},
		},

		question784{
			para784{"12345"},
			ans784{[]string{"12345"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 784------------------------\n")

	for _, q := range qs {
		_, p := q.ans784, q.para784
		fmt.Printf("【input】:%v       【output】:%v\n", p, letterCasePermutation1(p.one))
	}
	fmt.Printf("\n\n\n")
}
