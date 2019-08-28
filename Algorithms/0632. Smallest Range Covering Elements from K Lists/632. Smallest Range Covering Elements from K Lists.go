package leetcode

import (
	"math"
	"sort"
)

func smallestRange(nums [][]int) []int {
	numList, left, right, count, freqMap, res, length := []element{}, 0, -1, 0, map[int]int{}, make([]int, 2), math.MaxInt64
	for i, ns := range nums {
		for _, v := range ns {
			numList = append(numList, element{val: v, index: i})
		}
	}
	sort.Sort(SortByVal{numList})
	for left < len(numList) {
		if right+1 < len(numList) && count < len(nums) {
			right++
			if freqMap[numList[right].index] == 0 {
				count++
			}
			freqMap[numList[right].index]++
		} else {
			if count == len(nums) {
				if numList[right].val-numList[left].val < length {
					length = numList[right].val - numList[left].val
					res[0] = numList[left].val
					res[1] = numList[right].val
				}
			}
			freqMap[numList[left].index]--
			if freqMap[numList[left].index] == 0 {
				count--
			}
			left++
		}
	}
	return res
}

type element struct {
	val   int
	index int
}

type elements []element

// Len define
func (p elements) Len() int { return len(p) }

// Swap define
func (p elements) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// SortByVal define
type SortByVal struct{ elements }

// Less define
func (p SortByVal) Less(i, j int) bool {
	return p.elements[i].val < p.elements[j].val
}
