package leetcode

import (
	"sort"
)

func sortEvenOdd(nums []int) []int {
	odd, even, res := []int{}, []int{}, []int{}
	for index, v := range nums {
		if index%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	sort.Ints(even)
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))

	indexO, indexE := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			res = append(res, even[indexE])
			indexE++
		} else {
			res = append(res, odd[indexO])
			indexO++
		}
	}
	return res
}
