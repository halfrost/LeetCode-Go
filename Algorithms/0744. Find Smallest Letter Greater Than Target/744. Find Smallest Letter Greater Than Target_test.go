package leetcode

import (
	"fmt"
	"testing"
)

type question744 struct {
	para744
	ans744
}

// para 是参数
// one 代表第一个参数
type para744 struct {
	letters []byte
	target  byte
}

// ans 是答案
// one 代表第一个答案
type ans744 struct {
	one byte
}

func Test_Problem744(t *testing.T) {

	qs := []question744{

		question744{
			para744{[]byte{'c', 'f', 'j'}, 'a'},
			ans744{'c'},
		},

		question744{
			para744{[]byte{'c', 'f', 'j'}, 'c'},
			ans744{'f'},
		},

		question744{
			para744{[]byte{'c', 'f', 'j'}, 'd'},
			ans744{'f'},
		},

		question744{
			para744{[]byte{'c', 'f', 'j'}, 'g'},
			ans744{'j'},
		},

		question744{
			para744{[]byte{'c', 'f', 'j'}, 'j'},
			ans744{'c'},
		},

		question744{
			para744{[]byte{'c', 'f', 'j'}, 'k'},
			ans744{'c'},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 744------------------------\n")

	for _, q := range qs {
		_, p := q.ans744, q.para744
		fmt.Printf("【input】:%v       【output】:%v\n", p, nextGreatestLetter(p.letters, p.target))
	}
	fmt.Printf("\n\n\n")
}
