package leetcode

import (
	"fmt"
	"testing"
)

// intNI builds a NestedInteger holding a single integer.
func intNI(v int) *NestedInteger {
	n := &NestedInteger{}
	n.SetInteger(v)
	return n
}

// listNI builds a NestedInteger holding a nested list.
func listNI(elems ...*NestedInteger) *NestedInteger {
	n := &NestedInteger{}
	for _, e := range elems {
		n.Add(*e)
	}
	return n
}

func Test_Problem338(t *testing.T) {
	// Empty list: HasNext is false, Next returns -1.
	obj := Constructor([]*NestedInteger{})
	fmt.Printf("obj = %v\n", obj)
	if obj.HasNext() {
		t.Fatalf("expected HasNext to be false for empty list")
	}
	if got := obj.Next(); got != -1 {
		t.Fatalf("expected Next to return -1 for empty list, got %d", got)
	}

	// Nested input: [[1,1],2,[1,1]] should flatten to 1,1,2,1,1.
	// This exercises the nested-list branch (push) and the integer branch.
	nestedList := []*NestedInteger{
		listNI(intNI(1), intNI(1)),
		intNI(2),
		listNI(intNI(1), intNI(1)),
	}
	want := []int{1, 1, 2, 1, 1}
	obj2 := Constructor(nestedList)
	var got []int
	for obj2.HasNext() {
		got = append(got, obj2.Next())
	}
	fmt.Printf("flattened = %v\n", got)
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}

	// Verify GetInteger / GetList / IsInteger stub behavior.
	leaf := intNI(7)
	if !leaf.IsInteger() {
		t.Fatalf("expected leaf to be integer")
	}
	if leaf.GetInteger() != 7 {
		t.Fatalf("expected GetInteger 7, got %d", leaf.GetInteger())
	}
	branch := listNI(intNI(7))
	if branch.IsInteger() {
		t.Fatalf("expected branch to be a list")
	}
	if len(branch.GetList()) != 1 {
		t.Fatalf("expected GetList length 1, got %d", len(branch.GetList()))
	}
}
