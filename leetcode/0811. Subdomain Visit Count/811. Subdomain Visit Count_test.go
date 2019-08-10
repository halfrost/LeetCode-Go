package leetcode

import (
	"fmt"
	"testing"
)

type question811 struct {
	para811
	ans811
}

// para 是参数
// one 代表第一个参数
type para811 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans811 struct {
	one []string
}

func Test_Problem811(t *testing.T) {

	qs := []question811{

		question811{
			para811{[]string{"9001 discuss.leetcode.com"}},
			ans811{[]string{"mqe", "mqE", "mQe", "mQE", "Mqe", "MqE", "MQe", "MQE"}},
		},

		question811{
			para811{[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}},
			ans811{[]string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 811------------------------\n")

	for _, q := range qs {
		_, p := q.ans811, q.para811
		fmt.Printf("【input】:%v       【output】:%v\n", p, subdomainVisits(p.one))
	}
	fmt.Printf("\n\n\n")
}
