# [341. Flatten Nested List Iterator](https://leetcode.com/problems/flatten-nested-list-iterator/)


## 题目

Given a nested list of integers, implement an iterator to flatten it.

Each element is either an integer, or a list -- whose elements may also be integers or other lists.

**Example 1:**

```
Input:[[1,1],2,[1,1]]
Output:[1,1,2,1,1]
Explanation:By callingnext repeatedly untilhasNext returns false,
             the order of elements returned bynext should be:[1,1,2,1,1].
```

**Example 2:**

```
Input:[1,[4,[6]]]
Output:[1,4,6]
Explanation:By callingnext repeatedly untilhasNext returns false,
             the order of elements returned bynext should be:[1,4,6].

```

## 题目大意

给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。

## 解题思路

- 题目要求实现一个嵌套版的数组。可以用 `[]int` 实现，也可以用链表实现。笔者此处用链表实现。外层构造一个一维数组，一维数组内部每个元素是一个链表。额外还需要记录这个嵌套链表在原数组中的 `index` 索引。`Next()` 实现比较简单，取出对应的嵌套节点。`HasNext()` 方法则感觉嵌套节点里面的 `index` 信息判断是否还有 `next` 元素。

## 代码

```go
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    stack *list.List
}

type listIndex struct {
    nestedList []*NestedInteger
    index int
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
```