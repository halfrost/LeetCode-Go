package leetcode

import "container/list"

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return true }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return nil }

type NestedIterator struct {
	stack *list.List
}

type listIndex struct {
	nestedList []*NestedInteger
	index      int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := list.New()
	stack.PushBack(&listIndex{nestedList, 0})
	return &NestedIterator{stack}
}

func (this *NestedIterator) Next() int {
	if !this.HasNext() {
		return -1
	}
	last := this.stack.Back().Value.(*listIndex)
	nestedList, i := last.nestedList, last.index
	val := nestedList[i].GetInteger()
	last.index++
	return val
}

func (this *NestedIterator) HasNext() bool {
	stack := this.stack
	for stack.Len() > 0 {
		last := stack.Back().Value.(*listIndex)
		nestedList, i := last.nestedList, last.index
		if i >= len(nestedList) {
			stack.Remove(stack.Back())
		} else {
			val := nestedList[i]
			if val.IsInteger() {
				return true
			}
			last.index++
			stack.PushBack(&listIndex{val.GetList(), 0})
		}
	}
	return false
}
