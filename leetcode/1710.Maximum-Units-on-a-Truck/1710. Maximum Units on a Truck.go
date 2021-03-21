package leetcode

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	res := 0
	for i := 0; truckSize > 0 && i < len(boxTypes); i++ {
		if truckSize >= boxTypes[i][0] {
			truckSize -= boxTypes[i][0]
			res += (boxTypes[i][1] * boxTypes[i][0])
		} else {
			res += (truckSize * boxTypes[i][1])
			truckSize = 0
		}
	}
	return res
}
