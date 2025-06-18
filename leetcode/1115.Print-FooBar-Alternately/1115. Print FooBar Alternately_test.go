package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type question1115 struct {
	para1115
	ans1115
}

// para 是参数
// one 代表第一个参数
type para1115 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1115 struct {
	one string
}

func Test_Problem1115(t *testing.T) {

	qs := []question1115{

		{
			para1115{1},
			ans1115{"foobar"},
		},

		{
			para1115{2},
			ans1115{"foobarfoobar"},
		},

		{
			para1115{3},
			ans1115{"foobarfoobarfoobar"},
		},

		{
			para1115{4},
			ans1115{"foobarfoobarfoobarfoobar"},
		},

		{
			para1115{5},
			ans1115{"foobarfoobarfoobarfoobarfoobar"},
		},

		{
			para1115{6},
			ans1115{"foobarfoobarfoobarfoobarfoobarfoobar"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1115------------------------\n")

	for _, q := range qs {
		fooBarIns := NewFooBar(q.para1115.one)
		fmt.Printf("【input】:%v       【output】:", q.para1115.one)
		coin := rand.Intn(1)
		if coin == 0 {
			go fooBarIns.foo()
			go fooBarIns.bar()
		} else {
			go fooBarIns.bar()
			go fooBarIns.foo()
		}
		time.Sleep(time.Second)
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\n")
}
