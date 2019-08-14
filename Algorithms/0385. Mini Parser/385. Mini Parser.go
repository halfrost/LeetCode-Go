package leetcode

import (
	"fmt"
	"strconv"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

// NestedInteger define
type NestedInteger struct {
	Num  int
	List []*NestedInteger
}

// IsInteger define
func (n NestedInteger) IsInteger() bool {
	if n.List == nil {
		return true
	}
	return false
}

// GetInteger define
func (n NestedInteger) GetInteger() int {
	return n.Num
}

// SetInteger define
func (n *NestedInteger) SetInteger(value int) {
	n.Num = value
}

// Add define
func (n *NestedInteger) Add(elem NestedInteger) {
	n.List = append(n.List, &elem)
}

// GetList define
func (n NestedInteger) GetList() []*NestedInteger {
	return n.List
}

// Print define
func (n NestedInteger) Print() {
	if len(n.List) != 0 {
		for _, v := range n.List {
			if len(v.List) != 0 {
				v.Print()
				return
			}
			fmt.Printf("%v ", v.Num)
		}
	} else {
		fmt.Printf("%v ", n.Num)
	}
	fmt.Printf("\n")
}

func deserialize(s string) *NestedInteger {
	stack, cur := []*NestedInteger{}, &NestedInteger{}
	for i := 0; i < len(s); {
		switch {
		case isDigital(s[i]) || s[i] == '-':
			j := 0
			for j = i + 1; j < len(s) && isDigital(s[j]); j++ {
			}
			num, _ := strconv.Atoi(s[i:j])
			next := &NestedInteger{}
			next.SetInteger(num)
			if len(stack) > 0 {
				stack[len(stack)-1].List = append(stack[len(stack)-1].GetList(), next)
			} else {
				cur = next
			}
			i = j
		case s[i] == '[':
			next := &NestedInteger{}
			if len(stack) > 0 {
				stack[len(stack)-1].List = append(stack[len(stack)-1].GetList(), next)
			}
			stack = append(stack, next)
			i++
		case s[i] == ']':
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i++
		case s[i] == ',':
			i++
		}
	}
	return cur
}
