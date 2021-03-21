package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
	"sort"
)

// 解法一 线段树 SegmentTree
func createSortedArray(instructions []int) int {
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

// 解法二 树状数组 Binary Indexed Tree
func createSortedArray1(instructions []int) int {
	b := newBIT(make([]int, 100001))
	var res int
	cnt := map[int]int{}
	for i, n := range instructions {
		less := b.get(n - 1)
		greater := i - less - cnt[n]
		res = (res + min(less, greater)) % (1e9 + 7)
		b.update(n, 1)
		cnt[n]++
	}

	return res % (1e9 + 7)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type BIT struct {
	data []int
}

func newBIT(nums []int) *BIT {
	data := make([]int, len(nums)+1)
	b := &BIT{data}
	for i, n := range nums {
		b.update(i, n)
	}

	return b
}

func (b *BIT) update(i, num int) {
	i++
	for i < len(b.data) {
		b.data[i] += num
		i += (i & -i)
	}
}

func (b *BIT) get(i int) int {
	i++
	var sum int
	for i > 0 {
		sum += b.data[i]
		i -= (i & -i)
	}
	return sum
}
