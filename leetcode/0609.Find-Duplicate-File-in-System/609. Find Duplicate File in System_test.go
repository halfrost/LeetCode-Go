package leetcode

import (
	"fmt"
	"testing"
)

type question609 struct {
	para609
	ans609
}

// para 是参数
// one 代表第一个参数
type para609 struct {
	paths []string
}

// ans 是答案
// one 代表第一个答案
type ans609 struct {
	one [][]string
}

func Test_Problem609(t *testing.T) {

	qs := []question609{

		{
			para609{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}},
			ans609{[][]string{{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}}},
		},

		{
			para609{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"}},
			ans609{[][]string{{"root/a/2.txt", "root/c/d/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 609------------------------\n")

	for _, q := range qs {
		_, p := q.ans609, q.para609
		fmt.Printf("【input】:%v       【output】:%v\n", p, findDuplicate(p.paths))
	}
	fmt.Printf("\n\n\n")
}
