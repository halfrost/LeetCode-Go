package leetcode

import (
	"fmt"
	"testing"
)

type question385 struct {
	para385
	ans385
}

// para 是参数
// one 代表第一个参数
type para385 struct {
	n string
}

// ans 是答案
// one 代表第一个答案
type ans385 struct {
	one []int
}

func Test_Problem385(t *testing.T) {

	qs := []question385{

		{
			para385{"[[]]"},
			ans385{[]int{}},
		},

		{
			para385{"[]"},
			ans385{[]int{}},
		},

		{
			para385{"[-1]"},
			ans385{[]int{-1}},
		},

		{
			para385{"[123,[456,[789]]]"},
			ans385{[]int{123, 456, 789}},
		},

		{
			para385{"324"},
			ans385{[]int{324}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 385------------------------\n")

	for _, q := range qs {
		_, p := q.ans385, q.para385
		fmt.Printf("【input】:%v       【output】: \n", p)
		fmt.Printf("NestedInteger = ")
		ni := deserialize(p.n)
		ni.Print()

		// exercise NestedInteger interface methods
		if ni.IsInteger() {
			if ni.GetInteger() != ni.Num {
				t.Fatalf("GetInteger mismatch: got %v, want %v", ni.GetInteger(), ni.Num)
			}
		}
	}

	// exercise Add, IsInteger, GetInteger explicitly
	parent := &NestedInteger{}
	child := NestedInteger{}
	child.SetInteger(42)
	if !child.IsInteger() {
		t.Fatalf("expected child to be integer")
	}
	if child.GetInteger() != 42 {
		t.Fatalf("GetInteger mismatch: got %v, want 42", child.GetInteger())
	}
	parent.Add(child)
	if parent.IsInteger() {
		t.Fatalf("expected parent to be a list after Add")
	}
	if len(parent.GetList()) != 1 {
		t.Fatalf("expected list length 1, got %v", len(parent.GetList()))
	}
	fmt.Printf("\n\n\n")
}
