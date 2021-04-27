package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/template"
)

// 解法一 树状数组 Binary Indexed Tree
func createSortedArray(instructions []int) int {
	bit, res := template.BinaryIndexedTree{}, 0
	bit.Init(100001)
	for i, v := range instructions {
		less := bit.Query(v - 1)
		greater := i - bit.Query(v)
		res = (res + min(less, greater)) % (1e9 + 7)
		bit.Add(v, 1)
	}
	return res
}

// 解法二 线段树 SegmentTree
func createSortedArray1(instructions []int) int {
	if len(instructions) == 0 {
		return 0
	}
	st, res, mod := template.SegmentCountTree{}, 0, 1000000007
	numsMap, numsArray, tmpArray := discretization1649(instructions)
	// 初始化线段树，节点内的值都赋值为 0，即计数为 0
	st.Init(tmpArray, func(i, j int) int {
		return 0
	})
	for i := 0; i < len(instructions); i++ {
		strictlyLessThan := st.Query(0, numsMap[instructions[i]]-1)
		strictlyGreaterThan := st.Query(numsMap[instructions[i]]+1, numsArray[len(numsArray)-1])
		res = (res + min(strictlyLessThan, strictlyGreaterThan)) % mod
		st.UpdateCount(numsMap[instructions[i]])
	}
	return res
}

func discretization1649(instructions []int) (map[int]int, []int, []int) {
	tmpArray, numsArray, numsMap := []int{}, []int{}, map[int]int{}
	for i := 0; i < len(instructions); i++ {
		numsMap[instructions[i]] = instructions[i]
	}
	for _, v := range numsMap {
		numsArray = append(numsArray, v)
	}
	sort.Ints(numsArray)
	for i, num := range numsArray {
		numsMap[num] = i
	}
	for i := range numsArray {
		tmpArray = append(tmpArray, i)
	}
	return numsMap, numsArray, tmpArray
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
