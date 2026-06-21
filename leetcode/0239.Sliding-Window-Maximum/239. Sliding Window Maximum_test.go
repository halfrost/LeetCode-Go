package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question239 struct {
	para239
	ans239
}

// para 是参数
// one 代表第一个参数
type para239 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans239 struct {
	one []int
}

func Test_Problem239(t *testing.T) {

	qs := []question239{

		{
			para239{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			ans239{[]int{3, 3, 5, 5, 6, 7}},
		},
		{
			para239{[]int{}, 3},
			ans239{[]int{}},
		},
		{
			para239{[]int{1, 2}, 3},
			ans239{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 239------------------------\n")

	for _, q := range qs {
		a, p := q.ans239, q.para239
		got := maxSlidingWindow(p.one, p.k)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if len(got) != len(a.one) || (len(got) > 0 && !reflect.DeepEqual(got, a.one)) {
			t.Fatalf("maxSlidingWindow(%v, %d) = %v, want %v", p.one, p.k, got, a.one)
		}
		got1 := maxSlidingWindow1(p.one, p.k)
		if len(got1) != len(a.one) || (len(got1) > 0 && !reflect.DeepEqual(got1, a.one)) {
			t.Fatalf("maxSlidingWindow1(%v, %d) = %v, want %v", p.one, p.k, got1, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
