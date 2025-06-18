package leetcode

import (
	"fmt"
	"testing"
	"time"
)

type question1114 struct {
	para1114
	ans1114
}

// para 是参数
// one 代表第一个参数
type para1114 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1114 struct {
	one string
}

func Test_Problem1114(t *testing.T) {

	qs := []question1114{

		{
			para1114{[]int{1, 2, 3}},
			ans1114{"firstsecondthird"},
		},

		{
			para1114{[]int{1, 3, 2}},
			ans1114{"firstsecondthird"},
		},

		{
			para1114{[]int{2, 1, 3}},
			ans1114{"firstsecondthird"},
		},

		{
			para1114{[]int{2, 3, 1}},
			ans1114{"firstsecondthird"},
		},

		{
			para1114{[]int{3, 1, 2}},
			ans1114{"firstsecondthird"},
		},

		{
			para1114{[]int{3, 2, 1}},
			ans1114{"firstsecondthird"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1114------------------------\n")

	for _, q := range qs {
		_, p := q.ans1114, q.para1114
		fooIns := NewFoo()
		fmt.Printf("【input】:%v       【output】:", p)
		for _, v := range q.para1114.one {
			switch v {
			case 1:
				go fooIns.first()
			case 2:
				go fooIns.second()
			case 3:
				go fooIns.third()
			default:
				panic("problem does not support input other than '1, 2, 3'")
			}
		}
		time.Sleep(time.Second)
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\n")
}
